import { useState } from 'react';

export type Value = string | number;

interface UserEditableCellProps {
  value: string;
  updateUser: (value: Value) => Promise<boolean>;
}

export const UserEditableCell = ({
  value: currentValue,
  updateUser,
}: UserEditableCellProps) => {
  const [isEditing, setIsEditing] = useState(false);
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [draft, setDraft] = useState(currentValue);

  const startEditing = () => {
    setDraft(currentValue);
    setIsEditing(true);
  };

  const handleEnterPress = async (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === 'Enter') {
      e.preventDefault();
      handleSubmit();
    } else if (e.key === 'Escape') {
      setIsEditing(false);
    }
  };

  const handleSubmit = async () => {
    if (draft === currentValue) {
      setIsEditing(false);

      return;
    }

    setIsSubmitting(true);

    const isSuccess = await updateUser(draft);

    if (!isSuccess) {
      setDraft(currentValue);
    }

    setIsEditing(false);
    setIsSubmitting(false);
  };

  if (!isEditing) {
    return <span onClick={startEditing}>{currentValue}</span>;
  }

  return (
    <input
      className="border-2 border-gray-500 rounded-md disabled:opacity-50 disabled:pointer-events-none px-2 py-1 w-full sm:w-auto"
      type="text"
      id="name"
      name="name"
      value={String(draft)}
      onChange={(e) => setDraft(e.target.value)}
      onKeyDown={handleEnterPress}
      onBlur={handleSubmit}
      disabled={isSubmitting}
      autoFocus
    />
  );
};
