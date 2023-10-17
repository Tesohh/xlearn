import type { User } from '$lib/types';
import { backendUrl } from './const';

/* export const makeRequest = async (
	url: string,
	method: 'GET' | 'POST',
	body?: Object,
	cookie?: string
): Promise<{ error: boolean; content: Object }> => {
	if (method == 'GET') {
		let resp = await fetch(url, {
			headers: {
				Cookie: `jwt=${cookie}`
			}
		});

		console.log(resp);
	}

	let resp = await fetch(url, {
		method: 'POST',
		headers: {
			Cookie: `jwt=${cookie}`,
			'content-type': 'application/json; charset=utf-8'
		},
		body: JSON.stringify(body)
	});

	if (!resp.ok) return { error: true };

	console.log(resp.headers.getSetCookie());

	return { content: await resp.json() };
}; */

export const login = async (
	username: string,
	password: string
): Promise<{ error: boolean; cookie: string }> => {
	let resp = await fetch(`${backendUrl}/api/user/login`, {
		method: 'POST',
		headers: {
			'content-type': 'application/json; charset=utf-8'
		},
		body: JSON.stringify({ username: username, password: password })
	});

	if (!resp.ok) return { error: true, cookie: '' };

	return { error: false, cookie: resp.headers.getSetCookie()[0].replace('token=', '') };
};

export const cookieToUser = async (
	cookie: string
): Promise<{ error: boolean; user: User | null }> => {
	let resp = await fetch(`${backendUrl}/api/user/me`, {
		method: 'GET',
		headers: {
			Cookie: `token=${cookie}`
		}
	});

	if (!resp.ok) return { error: true, user: null };

	return { error: false, user: await resp.json() };
};

export const register = async (username: string, password: string) => {
	let resp = await fetch(`${backendUrl}/api/user/signup`, {
		method: 'POST',
		headers: {
			'content-type': 'application/json'
		},
		body: JSON.stringify({
			username: username,
			password: password
		})
	});

	if (!resp.ok) return { error: true };

	return { error: false };
};
