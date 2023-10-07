import { backendUrl } from '$lib/const';

export const handle = async ({ event, resolve }) => {
	const cookie = event.cookies.get('auth');

	if (cookie) {
		const resp = await fetch(`${backendUrl}/api/user/me`, {
			method: 'GET',
			headers: {
				Cookie: `jwt=${cookie}`
			}
		});

		if (!resp.ok) event.locals.user = null;
		else {
			let jsonResp = await resp.json();
			console.log(jsonResp);
		}
	} else {
		event.locals.user = null;
	}

	const response = await resolve(event);
	return response;
};
