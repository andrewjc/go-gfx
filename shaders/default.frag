#version 410 core

out vec4 FragColor;

in vec3 FragPos;
in vec3 Normal;

uniform vec3 objectColor;

void main()
{
    vec3 lightColor = vec3(1.0, 1.0, 1.0); // Uniform light color
    vec3 lightDirection = normalize(vec3(1.0, 1.0, 1.0)); // Uniform light direction

    // Calculate diffuse lighting
    float diffuse = max(dot(Normal, lightDirection), 0.0);

    // Calculate final color as the object color multiplied by the diffuse lighting factor
    vec3 finalColor = objectColor * diffuse;

    FragColor = vec4(finalColor, 1.0);
}
