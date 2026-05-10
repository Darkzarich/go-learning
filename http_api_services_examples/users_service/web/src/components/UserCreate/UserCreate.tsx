import { useState } from 'react';
import { createUser } from '../../api/users';
import { AxiosError } from 'axios';
import { Button } from '../Button/Button';

interface UserCreateProps {
  onCreated: () => void;
}

export const UserCreate = ({ onCreated }: UserCreateProps) => {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [error, setError] = useState('');
  const [isSubmitting, setIsSubmitting] = useState(false);

  const isNotFilled = name.trim() === '' || email.trim() === '';

  const handleSubmit = async (e: React.SubmitEvent<HTMLFormElement>) => {
    e.preventDefault();

    try {
      setError('');
      setIsSubmitting(true);

      await createUser({
        name,
        email,
      });

      setName('');
      setEmail('');
      setError('');

      onCreated();
    } catch (err) {
      if (err instanceof AxiosError) {
        console.log('error :>> ', err.response);
        setError(err.response?.data.error);
        return;
      }
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="flex flex-col gap-4 p-4">
      <div className="flex flex-col sm:flex-row items-start sm:items-center gap-4">
        <div className="flex flex-col sm:flex-row sm:items-center gap-2">
          <label htmlFor="name" className="font-bold min-w-14">
            Name
          </label>
          <input
            type="text"
            id="name"
            name="name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            className="ring-2 ring-gray-500 focus:border-transparent rounded-md px-2 py-1 w-full sm:w-auto outline-none focus-visible:ring-2 focus-visible:ring-blue-500"
          />
        </div>

        <div className="flex flex-col sm:flex-row sm:items-center gap-2">
          <label htmlFor="email" className="font-bold min-w-14">
            Email
          </label>
          <input
            type="email"
            id="email"
            name="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            className="ring-2 ring-gray-500 border-gray-500 focus:border-transparent rounded-md px-2 py-1 w-full sm:w-auto outline-none focus-visible:ring-2 focus-visible:ring-blue-500"
          />
        </div>

        <Button disabled={isSubmitting || isNotFilled}>Create</Button>
      </div>
      {error && <p className="text-red-500">{error}</p>}
    </form>
  );
};
