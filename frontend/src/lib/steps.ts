import { backendUrl } from './const';
import { parseStep, type Adventure, type Step } from '$lib/types';

export const getAdventureSteps = async (
	adventure: Adventure,
	cookie: string
): Promise<{ error: null | boolean; steps: null | Step[] }> => {
	let steps = new Array<Step>();

	adventure.steps.forEach(async (step) => {
		console.log(step);
		let resp;
		try {
			resp = await fetch(`${backendUrl}/api/step/@${step}`, {
				method: 'GET',
				headers: {
					Cookie: `token=${cookie}`
				}
			});
			console.log('REQUEST');
			console.log(await resp.json());
		} catch (err) {
			return { error: true, steps: null };
		}
		if (!resp.ok) return { error: true, steps: null };

		const json = await resp.json();

		const parsed = parseStep(json['step']);

		if (parsed.error) return { errror: true, steps: null };

		// @ts-ignore
		steps.push(parsed.step);
	});

	console.log(steps);

	return { error: false, steps: steps };
};
