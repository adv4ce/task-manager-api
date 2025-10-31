const API_URL = window.API_URL

async function handleFetch(res) {
	const text = await res.text()
	let data
	try {
		data = JSON.parse(text)
	} catch {
		data = { status: 'error', error: text || 'Unknown error' }
	}
	return data
}

export async function Create(title, description, priority) {
	const requestBody = { title, description, priority }
	const res = await fetch(API_URL, {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify(requestBody),
	})

	const responseData = await handleFetch(res)
	console.log('✅ Task created:', responseData)

	return {
		request: { method: 'POST /api/tasks', body: requestBody },
		response: responseData,
	}
}

export async function Get(id) {
	const res = await fetch(`${API_URL}/${id}`, {
		method: 'GET',
		headers: { 'Content-Type': 'application/json' },
	})
	const responseData = await handleFetch(res)
	console.log('✅ Task(s) received:', responseData)

	return {
		request: { method: `GET /api/tasks/${id}`, body: null },
		response: responseData,
	}
}

export async function Delete(id) {
	const res = await fetch(`${API_URL}/${id}`, {
		method: 'DELETE',
		headers: { 'Content-Type': 'application/json' },
	})
	const responseData = await handleFetch(res)
	console.log('✅ Task(s) deleted:', responseData)

	return {
		request: { method: `DELETE /api/tasks/${id}`, body: null },
		response: responseData,
	}
}

export async function Close(id) {
	const res = await fetch(`${API_URL}/${id}`, {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
	})
	const responseData = await handleFetch(res)
	console.log('✅ Task(s) closed:', responseData)

	return {
		request: { method: `POST /api/tasks/${id}`, body: null },
		response: responseData,
	}
}

export async function Update(id, title, description, priority) {
	const requestBody = { title, description, priority }
	const res = await fetch(`${API_URL}/${id}`, {
		method: 'PUT',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify(requestBody),
	})
	const responseData = await handleFetch(res)
	console.log('✅ Task updated:', responseData)

	return {
		request: { method: `PUT /api/tasks/${id}`, body: requestBody },
		response: responseData,
	}
}

export async function getAll() {
	const res = await fetch(API_URL, {
		method: 'GET',
	})

	const responseData = await handleFetch(res)
	console.log('✅ Task received:', responseData)

	return {
		request: { method: 'GET /api/tasks' },
		response: responseData,
	}
}
