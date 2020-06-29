#version 410

uniform mat4 camera;

in vec3 position;

void main(){
  gl_Position=vec4(position,1.);
}