#version 410 core

in vec3 vec_Color;

out vec4 color;

void main(){
  color=vec4(vec_Color,1);
}
