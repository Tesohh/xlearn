import { z } from 'zod';

export const User = z.object({
	display: z.string(),
	username: z.string(),
	coins: z.number(),
	role: z.number().optional(), // TODO sistemare dopo fix Tesohh
	joined_orgs: z.array(z.string()).optional() // TODO sistemare dopo fix Tesohh
});

export type UserType = z.infer<typeof User>;

export const Org = z.object({
	name: z.string(),
	tag: z.string(),
	secret: z.string().optional(),
	adventures: z.array(z.string()).optional()
});

export type OrgType = z.infer<typeof Org>;
