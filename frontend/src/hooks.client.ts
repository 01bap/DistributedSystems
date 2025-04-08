import type { HandleClientError } from '@sveltejs/kit';

export const handleError: HandleClientError = async ({ error, event, status, message }) => {
    const errorId = crypto.randomUUID() + "-fixed-" + status.toFixed();

    return {
        message: message + 'Whoops!',
        errorId
    };
};