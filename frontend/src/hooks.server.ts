import { cookieToUser } from '$lib/reqHandler';
import { authCookieName } from '$lib/const';

export const handle = async ({ event, resolve }) => {
	const cookie = event.cookies.get(authCookieName);

	if (cookie) {
		let result = await cookieToUser(cookie);
		// Check if the cookie is valid
		if (result?.error) event.locals.user = null;
		else {
			if (result.user) {
				event.locals.user = result.user;
			} else event.locals.user = null;
		}
	} else {
		event.locals.user = null;
	}

	const response = await resolve(event);
	return response;
};
