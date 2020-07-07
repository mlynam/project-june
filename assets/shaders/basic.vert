#version 410

uniform mat4 projection;
uniform mat4 camera;
uniform mat4 world;

layout(location=0)in vec3 position;

void main(){
  gl_Position=projection*camera*world*vec4(position,1.);
}