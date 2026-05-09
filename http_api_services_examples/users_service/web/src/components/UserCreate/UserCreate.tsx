import { useState } from 'react';
import { createUser } from '../../api/users';
import { AxiosError } from 'axios';

interface UserCreateProps {
  onCreated: () => void;
}

export const UserCreate = ({ onCreated }: UserCreateProps) => {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [error, setError] = useState('');
  const [isSubmitting, setIsSubmitting] = useState(false);

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
            className="border-2 border-gray-500 rounded-md px-2 py-1 w-full sm:w-auto"
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
            className="border-2 border-gray-500 rounded-md px-2 py-1 w-full sm:w-auto"
          />
        </div>
        <button
          className="bg-blue-500 hover:bg-blue-700 disabled:bg-gray-500 text-white font-bold py-2 px-4 rounded cursor-pointer"
          disabled={isSubmitting}
        >
          Create
        </button>
      </div>
      {error && <p className="text-red-500">{error}</p>}
    </form>
  );
};
