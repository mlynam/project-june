#version 410 core

uniform mat4 projection;
uniform mat4 view;
uniform mat4 world;

layout(location=0)in vec3 pos;
layout(location=1)in vec3 color;

out vec3 vec_Color;

void main(){
  gl_Position=projection*view*world*vec4(pos,1.);
  vec_Color=color;
}