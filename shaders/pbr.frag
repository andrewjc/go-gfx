#version 410 core

in vec3 fragNormal;
in vec3 fragPos;
in vec2 fragTexCoord;

out vec4 fragColor;

uniform vec3 camPos;
uniform vec3 albedo;
uniform float metallic;
uniform float roughness;

uniform sampler2D albedoMap;
uniform sampler2D normalMap;
uniform sampler2D metallicMap;
uniform sampler2D roughnessMap;
uniform sampler2D aoMap;

const float PI = 3.14159265359;

vec3 FresnelSchlick(float cosTheta, vec3 F0) {
    return F0 + (1.0 - F0) * pow(1.0 - cosTheta, 5.0);
}

float DistributionGGX(vec3 N, vec3 H, float roughness) {
    float a = roughness * roughness;
    float a2 = a * a;
    float NdotH = max(dot(N, H), 0.0);
    float NdotH2 = NdotH * NdotH;
    float d = (NdotH2 * (a2 - 1.0) + 1.0);
    return a2 / (PI * d * d);
}

float GeometrySchlickGGX(float NdotV, float roughness) {
    float r = (roughness + 1.0);
    float k = (r*r) / 8.0;

    float vis = NdotV / (NdotV * (1.0 - k) + k);
    return vis;
}

float GeometrySmith(vec3 N, vec3 V, vec3 L, float roughness) {
    float NdotV = max(dot(N, V), 0.0);
    float NdotL = max(dot(N, L), 0.0);
    float ggx2 = GeometrySchlickGGX(NdotV, roughness);
    float ggx1 = GeometrySchlickGGX(NdotL, roughness);
    return ggx1 * ggx2;
}

vec3 CookTorranceBRDF(vec3 N, vec3 V, vec3 L, vec3 albedo, float metallic, float roughness) {
    vec3 H = normalize(V + L);
    vec3 F0 = vec3(0.04);
    F0 = mix(F0, albedo, metallic);
    vec3 F = FresnelSchlick(max(dot(H, V), 0.0), F0);
    float D = DistributionGGX(N, H, roughness);
    float G = GeometrySmith(N, V, L, roughness);
    return (F * D * G) / (4.0 * max(dot(N, L), 0.0));
}

void main() {
    vec3 N = normalize(fragNormal);
    vec3 V = normalize(camPos - fragPos);
    vec3 albedoTex = texture(albedoMap, fragTexCoord).rgb;
    vec3 albedo = albedo * albedoTex;
    float metallicTex = texture(metallicMap, fragTexCoord).r;
    float metallic = metallic * metallicTex;
    float roughnessTex = texture(roughnessMap, fragTexCoord).r;
    float roughness = roughness * roughnessTex;
    vec3 L = normalize(vec3(-1.0, 0.5, -1.0));
    vec3 diffuse = albedo * (1.0 - metallic);
    vec3 specular = CookTorranceBRDF(N, V, L, albedo, metallic, roughness);
    vec3 ambient = vec3(0.03) * albedo;
    fragColor = vec4(diffuse + specular + ambient, 1.0);
}
