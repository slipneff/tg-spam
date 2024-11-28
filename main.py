import os
import asyncio
import uuid
from opentele.td import TDesktop
from opentele.api import CreateNewSession
from telethon.sessions import StringSession
import asyncpg

async def convert_tdata_to_session_string(tdata_folder):
    tdesk = TDesktop(tdata_folder)
    assert tdesk.isLoaded()

    client = await tdesk.ToTelethon("newSession.session", CreateNewSession)
    await client.connect()

    session_string = StringSession.save(client.session)
    await client.disconnect()

    return session_string

async def save_session_string_to_file(session_string, file_path):
    with open(file_path, 'w') as file:
        file.write(session_string)

async def save_session_info_to_db(session_id, file_path, db_connection):
    await db_connection.execute(
        "INSERT INTO sessions (id, path) VALUES ($1, $2)",
        session_id, file_path
    )

async def main():
    db_connection = await asyncpg.connect(
        user='tg',
        password='tg',
        database='tg',
        host='localhost:6432'
    )
    tdata_folder = os.path.join(os.path.dirname(__file__), 'tdata')

   
    for root, dirs, files in os.walk(tdata_folder):
        for file in files:
            tdata_path = os.path.join(root, file)
            try:
                session_string = await convert_tdata_to_session_string(tdata_path)

                session_id = str(uuid.uuid4())

                session_file_path = os.path.join(os.path.dirname(__file__), 'sessions', f"{session_id}.session")
                await save_session_string_to_file(session_string, session_file_path)

                await save_session_info_to_db(session_id, session_file_path, db_connection)

                print(f"Сохранена сессия для {tdata_path} в {session_file_path} с ID {session_id}")
            except Exception as e:
                print(f"Ошибка при обработке {tdata_path}: {e}")

    await db_connection.close()

asyncio.run(main())