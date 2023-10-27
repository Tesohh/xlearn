import { cookieToUser, parseUser } from '$lib/auth';
import { authCookieName } from '$lib/const';
import type { HandleFetch } from '@sveltejs/kit';

export const handle = async ({ event, resolve }) => {
	const cookie = event.cookies.get(authCookieName);

	if (cookie) {
		let result = await cookieToUser(cookie);
		// Check if the cookie is valid
		if (result?.error) {
			event.locals.user = null;
		} else {
			if (result.user) {
				const parsed = parseUser(result.user);
				if (!parsed.error || parsed.user) {
					event.locals.user = parsed.user;
				}
			} else event.locals.user = null;
		}
	} else {
		event.locals.user = null;
	}

	const response = await resolve(event);
	return response;
};

export const handleFetch = async ({ request, fetch }) => {
	if (request.url.startsWith('https://http.cat/')) {
		request.headers.set('origin', 'https://http.cat/');
	}

	return fetch(request);
};
