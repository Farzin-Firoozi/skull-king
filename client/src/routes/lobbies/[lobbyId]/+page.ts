import { redirect } from '@sveltejs/kit';
import AuthService from '../../../services/AuthService';
import ApiService from '../../../services/ApiService';
import { IntendedUrl } from '../../../constants';

/** @type {import('./$types').PageLoad} */
export async function load({ params }) {
	const auth = new AuthService();
	const user = auth.user();

	if (!user) {
		sessionStorage.setItem(IntendedUrl, window.location.href);
		throw redirect(302, '/login');
	} else if (!user.verified) {
		throw redirect(302, '/verify-email');
	}

	let ticketId = ""

	const apiService = new ApiService

	const response = await apiService.createTicket()

	if (response.status === 201) {
		const data = await response.json()
		ticketId = data.id
	} else {
		throw redirect(302, '/');
	}

	return {
		lobbyId: params.lobbyId,
		auth: user,
		ticketId: ticketId,
	}
}

export const ssr = false;
