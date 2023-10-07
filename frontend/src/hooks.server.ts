export const handle = async ({ event, resolve }) => {
	const cookie = event.cookies.get('auth');

	if (cookie) {
		event.locals.user = {
			username: 'test',
			level: 1
		};
	} else {
		event.locals.user = null;
	}

	const response = await resolve(event);
	return response;
};
