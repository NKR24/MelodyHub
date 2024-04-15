import axios from 'axios';
import type { User, UserID } from '../types';

const id = `98654381-405f-491d-ac37-687d96807e5a`;

export async function createUser(body: User) {
	await axios.post('http://localhost:4000/users', body);
}

export async function getUserId(): Promise<UserID[]> {
	const { data } = await axios.get(`http://localhost:4000/users/${id}`);
	return data;
}
