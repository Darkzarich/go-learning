import { ApiClient } from './client';

export interface CreateUserRequest {
  name: string;
  email: string;
}

export interface User {
  id: number;
  name: string;
  email: string;
  active: boolean;
  last_login: string;
  created_at: string;
}

export const getUsers = async (): Promise<User[]> => {
  const response = await ApiClient.get('/users');

  return response.data;
};

export const getUser = async (id: number): Promise<User> => {
  const response = await ApiClient.get(`/users/${id}`);

  return response.data;
};

export const createUser = async (user: CreateUserRequest) => {
  const response = await ApiClient.post('/users', user);

  return response.data;
};

export const updateUser = async (
  id: number,
  user: CreateUserRequest,
): Promise<User> => {
  const response = await ApiClient.put(`/users/${id}`, user);

  return response.data;
};

export const deleteUser = async (id: number) => {
  const response = await ApiClient.delete(`/users/${id}`);

  return response.data;
};
