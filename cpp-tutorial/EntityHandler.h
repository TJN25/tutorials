#include <memory>
#include "Entity.h"


class EntityHandler
{
private:
	int m_max_entities;
	// std::unique_ptr<Player[]> m_EntityArray;
	int m_EntityIndex[100];
	Player* m_EntityArray[100];
	static int m_entity_counter;
	void ReorderArray(int (&arr)[100], int a, int b);
public:
	EntityHandler(int max_entities);
	void PrintIndex();
	void MoveEntities();
};
