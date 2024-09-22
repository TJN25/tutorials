from sys import maxsize
import threading
import queue
import time
import sqlite3

class Worker(threading.Thread):
    def __init__(self, q: queue.Queue, *args, **kwargs):
        self.q: queue.Queue = q
        super().__init__(*args, **kwargs)

    def run(self):
        with sqlite3.connect(f'test.db') as conn:
            self.create_db(conn)
            while True:
                try:
                    work = self.q.get(timeout=1.5)  # 3s timeout
                except queue.Empty:
                    return

                self.add_to_db(conn, work)

                self.q.task_done()

    def create_db(self, conn: sqlite3.Connection):
        print(sqlite3.sqlite_version)
        sql_create_table = [ 
                """CREATE TABLE IF NOT EXISTS test (
                        id INT PRIMARY KEY, 
                        value INT,
                        constant INT
                );"""]
        cursor = conn.cursor()
        for statement in sql_create_table:
            cursor.execute(statement)

    def add_to_db(self, conn: sqlite3.Connection, value: int):
        time.sleep(0.5)
        cur = conn.cursor()
        sql = ''' INSERT INTO test(id,value,constant)
                  VALUES(?,?,?) '''
        data = (value * 100, value, 42)
        cur.execute(sql, data)
        conn.commit()
        return cur.lastrowid


q = queue.Queue(maxsize=500)
for _ in range(100):
    Worker(q).start()
q.join() 

for i in range(100, 2000):
    print(f'Put: {i}')
    q.put(i)
