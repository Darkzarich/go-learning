import { deleteUser } from '../../api/users';
import { useState, type ComponentProps } from 'react';
import { Button } from '../Button/Button';

interface UserDeleteBtnProps extends Omit<ComponentProps<'button'>, 'id'> {
  id: number;
  onDeleted: () => void;
}

export const UserDeleteBtn = ({
  id,
  onDeleted,
  ...props
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
    <Button
      variant="ghost-danger"
      onClick={() => handleDeleteUser(id)}
      disabled={isDeleting}
      {...props}
    >
      Delete
    </Button>
  );
};
