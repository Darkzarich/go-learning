import { useEffect, useState } from 'react';
import { type User, getUsers, updateUser } from './api/users';
import './App.css';
import {
  UserList,
  type UserUpdatableFields,
} from './components/UserList/UserList';
import { UserCreate } from './components/UserCreate/UserCreate';
import type { Value } from './components/UserEditableCell/UserEditableCell';
import { Button } from './components/Button/Button';

function App() {
  const [users, setUsers] = useState<User[]>([]);
  const [isFetching, setIsFetching] = useState(true);

  const fetchUsers = async () => {
    setIsFetching(true);

    const users = await getUsers();

    setUsers(users);
    setIsFetching(false);
  };

  const handleUpdateUser = async (
    id: number,
    field: UserUpdatableFields,
    value: Value,
  ) => {
    setIsFetching(true);

    const userToUpdate = users.find((user) => user.id === id);

    if (!userToUpdate) {
      setIsFetching(false);
      return false;
    }

    try {
      const updatedUser = {
        ...userToUpdate,
        [field]: value,
      };

      const newUsers = users.map((user) => {
        if (user.id === id) {
          return updatedUser;
        }

        return user;
      });

      await updateUser(id, updatedUser);

      setUsers(newUsers);

      return true;
    } catch (err) {
      console.log('error :>> ', err);

      return false;
    } finally {
      setIsFetching(false);
    }
  };

  useEffect(() => {
    let ignore = false;

    getUsers().then((users) => {
      if (!ignore) {
        setIsFetching(false);
        setUsers(users);
      }
    });

    return () => {
      ignore = true;
    };
  }, []);

  return (
    <div className="w-full max-w-5xl mx-auto px-4 py-6">
      <div className="flex flex-col gap-6">
        <div className="flex flex-col sm:flex-row sm:items-center gap-4">
          <h1 className="text-2xl font-bold">Users</h1>
          <Button onClick={fetchUsers} disabled={isFetching}>
            Refresh
          </Button>
        </div>

        <UserCreate onCreated={fetchUsers} />
        <UserList
          users={users}
          onDeleted={fetchUsers}
          onUpdate={handleUpdateUser}
        />
      </div>
    </div>
  );
}

export default App;
