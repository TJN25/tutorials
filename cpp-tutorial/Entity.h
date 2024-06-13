#pragma once

class Entity 
{
private:
	void UpdateSpeed(float xa, float ya);
public:
	Entity();
	~Entity();
	float x, y;
	float speed_x;
	float speed_y;
	float health;
	bool alive;
	int knockback_multiplier;
	int entity_id;
	static int EntityCounter;
	static float MAX_SPEED;
	void Move(float xa, float ya);
	void Hit(float hForce, float vForce);
	void UpdateHealth(float damage);
	void OutOfBounds();
	void Print();
};

class Player : public Entity
{
private:
	const char* m_name;
public:
	Player();
	~Player();
	void PrintName();
};
