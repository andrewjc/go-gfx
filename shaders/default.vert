#version 410 core

layout (location = 0) in vec3 position;
layout (location = 1) in vec3 color;

uniform mat4 projection;
uniform mat4 view;
uniform mat4 model;

out vec3 fragColor;

void main()
{
    gl_Position = projection * view * model * vec4(position, 1.0);
    gl_Position /= gl_Position.w;  // divide by w to get NDC
    fragColor = color;
}
