#version 410 core

in vec3 fragColor;

out vec4 fragColorOut;

void main()
{
    fragColorOut = vec4(fragColor, 1.0);
}

