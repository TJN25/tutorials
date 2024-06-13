#include <iostream>
#include "MyMath.h"
#include "Entity.h"

int Entity::EntityCounter = 0;
float Entity::MAX_SPEED = 5.0f;

Entity::Entity()
{
	knockback_multiplier = 10.0f;
	alive = true;
	entity_id = Entity::EntityCounter;
	Entity::EntityCounter += 1;
}
Entity::~Entity()
{
	Entity::EntityCounter -= 1;
}
void Entity::Move(float xa, float ya)
{
	UpdateSpeed(xa, ya);
	x += speed_x;
	y += speed_y;
	OutOfBounds();
}
void Entity::UpdateSpeed(float xa, float ya)
{
	speed_x += xa;
	speed_y += ya;
	if (speed_x > MAX_SPEED) speed_x = MAX_SPEED;
	if (speed_y > MAX_SPEED) speed_y = MAX_SPEED;
	if (speed_x < -MAX_SPEED) speed_x = -MAX_SPEED;
	if (speed_y < -MAX_SPEED) speed_y = -MAX_SPEED;
}

void Entity::Print()
{
	std::cout << x << ", " << y << std::endl;
}

void Entity::Hit(float hForce, float vForce)
{
	speed_x = speed_x/knockback_multiplier;
	speed_y = speed_y/knockback_multiplier;
	x += hForce/knockback_multiplier;
	y += vForce/knockback_multiplier;
	UpdateHealth(Pythagoras(hForce, vForce));
}

void Entity::UpdateHealth(float damage)
{
	std::cout << "Oof! " << damage << " health lost!" << health - damage << " remaining." << std::endl;
	health -= damage;
	if (health <= 0)
	{
		health = 0;
		alive = false;
		Entity::EntityCounter -= 1;
	}
}

void Entity::OutOfBounds() 
{
	if (x > 20)
		Hit(-100, 0);
	if (x < -20)
		Hit(100, 0);
	if (y > 20)
		Hit(0, 100);
	if (y < -20)
		Hit(0, 100);
}
