import { format } from 'date-fns';
import { type User } from '../../api/users';
import {
  UserEditableCell,
  type Value,
} from '../UserEditableCell/UserEditableCell';
import { UserDeleteBtn } from './UserDeleteBtn';

export type UserUpdatableFields = 'name' | 'email';

interface UserListProps {
  users: User[];
  onUpdate: (
    id: number,
    field: UserUpdatableFields,
    value: Value,
  ) => Promise<boolean>;
}

interface UserListProps {
  onDeleted: () => void;
}

const tableDateFormat = 'yyyy-MM-dd HH:mm';

export const UserList = ({ users, onDeleted, onUpdate }: UserListProps) => {
  return (
    <div className="overflow-x-auto w-full">
      <table className="w-full table-auto">
        <thead>
          <tr className="text-left">
            <th className="p-2 whitespace-nowrap">Name</th>
            <th className="p-2 whitespace-nowrap">Email</th>
            <th className="p-2 whitespace-nowrap">Active</th>
            <th className="p-2 whitespace-nowrap">Last login</th>
            <th className="p-2 whitespace-nowrap">Created at</th>
            <th className="p-2"></th>
          </tr>
        </thead>
        <tbody>
          {users.map((user) => (
            <tr key={user.id}>
              <td className="p-2">
                {
                  <UserEditableCell
                    value={user.name}
                    updateUser={(value) => onUpdate(user.id, 'name', value)}
                  />
                }
              </td>
              <td className="p-2">
                {
                  <UserEditableCell
                    value={user.email}
                    updateUser={(value) => onUpdate(user.id, 'email', value)}
                  />
                }
              </td>
              <td className="p-2">{user.active ? 'Yes' : 'No'}</td>
              <td className="p-2">
                {user.last_login
                  ? format(user.last_login, tableDateFormat)
                  : ''}
              </td>
              <td className="p-2">
                {user.created_at
                  ? format(user.created_at, tableDateFormat)
                  : ''}
              </td>
              <td className="p-2">
                <UserDeleteBtn
                  id={user.id}
                  onDeleted={onDeleted}
                  className="whitespace-nowrap"
                />
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};
