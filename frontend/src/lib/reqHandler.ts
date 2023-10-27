import { User, type UserType } from '$lib/types';
import { redirect } from '@sveltejs/kit';
import { backendUrl } from './const';
import type { ZodObject, output, z } from 'zod';

export const login = async (
	username: string,
	password: string
): Promise<{ error: boolean; cookie: string | null }> => {
	let resp: Response;

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

	if (!resp?.ok) return { error: true, cookie: null };

	return { error: false, cookie: resp.headers.getSetCookie()[0].replace('token=', '') };
};

export const cookieToUser = async (
	cookie: string
): Promise<{ error: boolean; user: UserType | null }> => {
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

	const parsed = parseUser(await resp.json());
	if (parsed.error || parsed.user == null) {
		return { error: true, user: null };
	}

	return { error: false, user: parsed.user };
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

	return { error: true };
};

export const parseUser = (userData: Object): { user: UserType | null; error: boolean } => {
	const parsed = User.safeParse(userData);

	if (!parsed.success) {
		console.log('ERROR');
		return { error: true, user: null };
	}
	return { user: parsed.data, error: false };
};
