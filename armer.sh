#!/bin/bash

debug=false

# Green tick emoji
success="\xE2\x9C\x85"
# Warning emoji
warning="\xE2\x9D\x97"
# Red cross mark emoji
failed="\xE2\x9D\x8C"

# Run the command and store the output in a variable
images=$(kubectl get pods --all-namespaces -o jsonpath="{.items[*].spec['initContainers', 'containers'][*].image}" |
    tr -s '[[:space:]]' '\n' |
    uniq)

echo -e "Legends:\n$success - Supports arm64, $failed - Do not support arm64, $warning - Some error occurred"
echo -e "------------------------------------------------------------------------\n"
if $debug; then
    echo -e "Debug mode is on.\n"
fi

# Process each line of the output
echo "$images" | while read line; do
    image=$line
    # If it is not a URL, append docker.io
    if [[ ! "$line" =~ ([a-zA-Z0-9\.\-]+\.[a-zA-Z0-9]+)\/([a-zA-Z0-9\.\-]+) ]]; then
        # Check if the line contains a slash
        if [[ $image == *"/"* ]]; then
            # If it does, prepend "docker.io/"
            image="docker.io/$image"
        else
            # If it doesn't, prepend "docker.io/library/"
            image="docker.io/library/$image"
        fi
    fi

    if $debug; then
        echo "Original - $line, Modified - $image"
        arch=$(skopeo inspect --override-os linux --no-tags --format "{{ .Architecture }}" --override-arch arm64 docker://$image)
    else
        arch=$(skopeo inspect --override-os linux --no-tags --format "{{ .Architecture }}" --override-arch arm64 docker://$image 2>/dev/null)
    fi

    # Get the exit code of the command
    exit_code=$?
    if [ $exit_code -ne 0 ]; then
        result=$warning
    else
        if [[ $arch == "arm64" ]]; then
            result=$success
        else
            result=$failed
        fi
    fi

    echo -e "$line $result\n"

done
