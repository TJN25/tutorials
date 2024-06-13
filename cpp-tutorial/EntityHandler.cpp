#include <iostream>
#include "Entity.h"
#include "EntityHandler.h"

EntityHandler::EntityHandler(int max_entities)
{
	m_max_entities = max_entities;
	if (max_entities > 100)
		m_max_entities = 100;

	for (int i = 0; i < m_max_entities; i++)
	{
		m_EntityIndex[i] = i;
		m_EntityArray[i] = new Player();
	}
	
}

void EntityHandler::MoveEntities()
{
	for (int i = 0; i < Entity::EntityCounter; i++)
	{
		int index = m_EntityIndex[i];
		m_EntityArray[index]->UpdateHealth(5 * (10 - i));
		if ( m_EntityArray[index]->alive == false )
		{
			ReorderArray(m_EntityIndex, i, Entity::EntityCounter);
		}
	}
}

void EntityHandler::PrintIndex()
{
	for(int i = 0; i < m_max_entities; i++){
		std::cout << "Index is: " << m_EntityIndex[i] << std::endl;
	}
}

void EntityHandler::ReorderArray(int (&arr)[100], int a, int b)
{
	int tmp_a = arr[a];
	int tmp_b = arr[b];
	arr[a] = tmp_b;
	arr[b] = tmp_a;
}

