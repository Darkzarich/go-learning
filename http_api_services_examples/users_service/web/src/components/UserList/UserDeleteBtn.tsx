import { deleteUser } from '../../api/users';
import { useState } from 'react';

interface UserDeleteBtnProps {
  id: number;
  onDeleted: () => void;
  className?: string;
}

export const UserDeleteBtn = ({
  id,
  onDeleted,
  className,
}: UserDeleteBtnProps) => {
  const [isDeleting, setIsDeleting] = useState(false);

  const handleDeleteUser = async (id: number) => {
    setIsDeleting(true);

    try {
      await deleteUser(id);
      onDeleted();
    } catch (err) {
      console.log('error :>> ', err);
    } finally {
      setIsDeleting(false);
    }
  };

  return (
    <button
      className={`${className} bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded`}
      onClick={() => handleDeleteUser(id)}
      disabled={isDeleting}
    >
      Delete
    </button>
  );
};
