#include <algorithm>
#include <cstddef>
#include <iostream>
#include <fstream>


void process_line_f(std::string& line, int ignore_count, char s) 
{
	if (line[0] != '>') 
		return;

	size_t replace_idx; 
	for (size_t idx = 1; idx < line.size(); idx++) 
	{	
		if (line[idx] != '[') 
			continue;;
		
		if (ignore_count == 0)
			line[idx] = '\t';
		ignore_count--;
	}
}

void process_line_e(std::string& line, int ignore_count, char s) 
{
	if (line[0] != '>') 
		return;

	size_t replace_idx; 
	for (size_t idx = line.size(); idx > 0; idx--) 
	{
		if (line[idx] != s) 
			continue;;
		
		if (ignore_count == 0)
			line[idx] = '\t';
		ignore_count--;

	}
}

void loop_file(std::string file_name, int position, int ignore_count, char s) 
{
	std::string line;
	std::ifstream f (file_name);

	if (f.is_open())
	{
		while ( getline (f,line) )
		{
			if (position == 0)
				process_line_f(line, ignore_count, s);
			else
				process_line_e(line, ignore_count, s);
			std::replace(line.begin(), line.end(), ' ', '*');
			std::cout << line << '\n';
		}
		f.close();
	}

	else std::cout << "Unable to open file"; 
}

int main(int argc, char *argv[]) 
{	
	std::string my_file = "test.faa";
	loop_file(my_file, 0, 0, '[');
	return 0;

}
