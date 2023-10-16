import { backendUrl } from './const';

export const makeRequest = async (
	url: string,
	method: 'GET' | 'POST',
	body?: Object,
	cookie?: string
): Promise<{ error?: boolean; content?: Object }> => {
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
};

export const login = async (
	username: string,
	password: string
): Promise<{ error?: boolean; cookie?: string[] }> => {
	let resp = await fetch(`${backendUrl}/api/user/login`, {
		method: 'POST',
		headers: {
			'content-type': 'application/json; charset=utf-8'
		},
		body: JSON.stringify({ username: username, password: password })
	});

	if (!resp.ok) return { error: true };

	return { cookie: resp.headers.getSetCookie() };
};
