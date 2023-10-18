import type { User } from '$lib/types';
import { redirect } from '@sveltejs/kit';
import { backendUrl } from './const';

export const login = async (
	username: string,
	password: string
): Promise<{ error: boolean; cookie: string }> => {
	let resp;

	try {
		resp = await fetch(`${backendUrl}/api/user/login`, {
			method: 'POST',
			headers: {
				'content-type': 'application/json; charset=utf-8'
			},
			body: JSON.stringify({ username: username, password: password })
		});
	} catch (err) {
		throw redirect(303, '/apiError');
	}

	if (!resp?.ok) return { error: true, cookie: '' };

	return { error: false, cookie: resp.headers.getSetCookie()[0].replace('token=', '') };
};

export const cookieToUser = async (
	cookie: string
): Promise<{ error: boolean; user: User | null }> => {
	let resp;

	try {
		resp = await fetch(`${backendUrl}/api/user/me`, {
			method: 'GET',
			headers: {
				Cookie: `token=${cookie}`
			}
		});
	} catch (err) {
		throw redirect(303, '/apiError');
	}

	if (!resp?.ok) return { error: true, user: null };

	return { error: false, user: await resp.json() };
};

export const register = async (username: string, password: string): Promise<{ error: boolean }> => {
	let resp;
	try {
		resp = await fetch(`${backendUrl}/api/user/signup`, {
			method: 'POST',
			headers: {
				'content-type': 'application/json'
			},
			body: JSON.stringify({
				username: username,
				password: password
			})
		});
	} catch (err) {
		throw redirect(303, '/apiError');
	}

	if (!resp.ok) return { error: true };

	return { error: false };
};
