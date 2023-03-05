#version 410 core

layout (location = 0) in vec3 aPos;
layout (location = 1) in vec3 aNormal;
layout (location = 2) in vec2 aTexCoord;

layout (location = 1) in vec3 color;

uniform mat4 projection;
uniform mat4 view;
uniform mat4 model;

out vec3 fragColor;

void main()
{
    gl_Position = projection * view * model * vec4(aPos, 1.0);

    // update for ndc coordinates
    gl_Position = gl_Position / gl_Position.w;

    fragColor = color;
}
