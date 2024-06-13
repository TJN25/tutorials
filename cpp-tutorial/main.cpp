#include <cstdlib>
#include <iostream>
#include "EntityHandler.h"



int main(int argc, char *argv[]) 
{	
	EntityHandler handler(3);
	std::cout << "Number of instances: " << Entity::EntityCounter << std::endl;
	while (Entity::EntityCounter > 0 )
	{
		std::cout << "Number of entites: " << Entity::EntityCounter << std::endl; 
		handler.MoveEntities();
	}
	// srand( (unsigned)time( NULL ) );
	// // Player* players[5];
	// std::array<Player*, 5> players;
	// for (int i = 0; i < 5; i++)
	// {
	// 	players[i] = new Player("tmp");
	// }
	// while (Entity::EntityCounter > 0)
	// {
	// 	for (int i = 0; i < 5; i++)
	// 	{
	// 		if (players[i]->alive == false) continue;
	// 		players[i]->Move(RandomNumberBetween(-1, 1), RandomNumberBetween(-1, 1));
	// 		players[i]->Print();
	// 	}
	// }
	// std::cout << "No entities left" << std::endl;
}
