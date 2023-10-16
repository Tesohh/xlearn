import { backendUrl } from '$lib/const';
import { makeRequest } from '$lib/reqHandler';

export const handle = async ({ event, resolve }) => {
	const cookie = event.cookies.get('auth');

	if (cookie) {
		let resp = await makeRequest(`${backendUrl}/api/user/me`, 'GET', {}, cookie);

		if (resp.error) event.locals.user = null;
		else {
			event.locals.user = {
				username: 'asa',
				level: 2
			};
		}
	} else {
		event.locals.user = null;
	}

	const response = await resolve(event);
	return response;
};
