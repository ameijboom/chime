import { redirect, error } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';
import { checkGuildMembershipAndPermissions } from '../auth';

export const load: LayoutServerLoad = async (event) => {
	const session = await event.locals.auth();

	if (!session?.user && !event.url.pathname.startsWith('/auth')) {
		throw redirect(303, '/auth/signin');
	}

	if (session?.user && session.accessToken && session.userId) {
		if (!(await checkGuildMembershipAndPermissions(session.accessToken))) {
			throw error(403, 'Only the Moderator team is allowed to access the Chime Dashboard.');
		}
	}

	return {
		session
	};
};
