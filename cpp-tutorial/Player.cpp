#include <iostream>
#include "Entity.h"


Player::Player()
{
	const char* name = "tmp";
	m_name = name;
	health = 300.0f;
	x = 0.0f;
	y = 0.0f;
	speed_x = 0.0f;
	speed_y = 0.0f;
}

Player::~Player()
{
	std::cout << "Player: " << m_name << " destroyed." << std::endl;
}

void Player::PrintName()
{
	std::cout << m_name << std::endl;
}
