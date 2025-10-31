import { Create, Get, Delete, Close, Update } from './api.js'
import { loadTasks } from './loadTasks.js'

async function loadTemplate(templateFile, targetSelector) {
	try {
		const response = await fetch(`/static/templates/${templateFile}`)
		const html = await response.text()
		document.querySelector(targetSelector).innerHTML = html

		initFormValidation()
	} catch (error) {
		console.error('Error loading template:', error)
	}
}

document.addEventListener('DOMContentLoaded', () => {
	loadTasks()
})

function switchTemplate(templateFile) {
	loadTemplate(templateFile, '#tabContent')
}

document.addEventListener('DOMContentLoaded', () => {
	loadTemplate('create-form.html', '#tabContent')

	const responseElement = document.getElementById('response-text')
	const requestElement = document.getElementById('requests-text')
	const status = document.querySelector('.response-status')

	document.getElementById('btnradio1').addEventListener('change', function () {
		if (this.checked) {
			switchTemplate('create-form.html')
			responseElement.innerHTML = ''
			requestElement.innerHTML = ''
			status.classList.add('status-unvisible')
		}
	})

	document.getElementById('btnradio2').addEventListener('change', function () {
		if (this.checked) {
			switchTemplate('get-form.html')
			responseElement.innerHTML = ''
			requestElement.innerHTML = ''
			status.classList.add('status-unvisible')
		}
	})

	document.getElementById('btnradio3').addEventListener('change', function () {
		if (this.checked) {
			switchTemplate('update.html')
			responseElement.innerHTML = ''
			requestElement.innerHTML = ''
			status.classList.add('status-unvisible')
		}
	})

	document.getElementById('btnradio4').addEventListener('change', function () {
		if (this.checked) {
			switchTemplate('delete.html')
			responseElement.innerHTML = ''
			requestElement.innerHTML = ''
			status.classList.add('status-unvisible')
		}
	})

	document.getElementById('btnradio5').addEventListener('change', function () {
		if (this.checked) {
			switchTemplate('close.html')
			responseElement.innerHTML = ''
			requestElement.innerHTML = ''
			status.classList.add('status-unvisible')
		}
	})
})

function validateForm() {
	const container = document.querySelector('#tabContent')
	const fields = container.querySelectorAll('input, textarea, select')
	const sendButton = document.getElementById('send-btn')

	let allFilled = true

	fields.forEach(field => {
		if (field.tagName === 'SELECT') return

		if (!field.value.trim()) {
			allFilled = false
		}
	})

	sendButton.disabled = !allFilled
	sendButton.classList.toggle('disabled', !allFilled)
}

function allowOnlyNumbers(input) {
	input.value = input.value.replace(/\D/g, '')
}

function initFormValidation() {
	const container = document.querySelector('#tabContent')
	const fields = container.querySelectorAll('input, textarea, select')

	fields.forEach(field => {
		if (field.id && field.id.includes('-id')) {
			field.addEventListener('input', () => {
				allowOnlyNumbers(field)
				validateForm()
			})
		} else {
			field.addEventListener('input', validateForm)
		}

		field.addEventListener('change', validateForm)
	})

	validateForm()
}

document.getElementById('send-btn').addEventListener('click', handleSend)

function getActiveOperation() {
	const active = document.querySelector('input[name="btnradio"]:checked')
	return active ? active.id : null
}

async function handleCreate() {
	const title = document.getElementById('create-title').value.trim()
	const description = document.getElementById('create-description').value.trim()
	const priority = document.querySelector('.custom-select-task').value

	const responseElement = document.getElementById('response-text')
	const requestElement = document.getElementById('requests-text')
	const status = document.querySelector('.response-status')
	const statusText = document.querySelector('.response-status-text')

	try {
		const data = await Create(title, description, priority)

		status.classList.remove(
			'status-unvisible',
			'status-success',
			'status-error'
		)
		if (data.response && data.response.status === 'successful') {
			status.classList.add('status-success', 'status-visible')
			statusText.textContent = 'Success'
		} else {
			status.classList.add('status-error', 'status-visible')
			statusText.textContent = 'Error'
		}

		requestElement.innerHTML = `<pre style="margin:0;font-family:inherit;">${
			data.request.method
		} ${JSON.stringify(data.request.body, null, 2)}</pre>`
		responseElement.innerHTML = `<pre style="margin:0;font-family:inherit;">${JSON.stringify(
			data.response,
			null,
			2
		)}</pre>`
	} catch (error) {
		status.classList.remove(
			'status-unvisible',
			'status-success',
			'status-error'
		)
		status.classList.add('status-error', 'status-visible')
		statusText.textContent = 'Error'

		responseElement.innerHTML = `<pre style="margin:0;font-family:inherit;">Error: ${error.message}</pre>`
		requestElement.innerHTML = ''
	}
}

async function handleGet() {
	const id = document.getElementById('get-id').value.trim()
	const responseElement = document.getElementById('response-text')
	const requestElement = document.getElementById('requests-text')
	const status = document.querySelector('.response-status')
	const statusText = document.querySelector('.response-status-text')

	try {
		const data = await Get(id)

		status.classList.remove(
			'status-unvisible',
			'status-success',
			'status-error'
		)
		if (data.response && data.response.status === 'successful') {
			status.classList.add('status-success', 'status-visible')
			statusText.textContent = 'Success'
		} else {
			status.classList.add('status-error', 'status-visible')
			statusText.textContent = 'Error'
		}

		requestElement.innerHTML = `<pre style="margin:0;font-family:inherit;">${data.request.method}</pre>`
		responseElement.innerHTML = `<pre style="margin:0;font-family:inherit;">${JSON.stringify(
			data.response,
			null,
			2
		)}</pre>`
	} catch (error) {
		status.classList.remove(
			'status-unvisible',
			'status-success',
			'status-error'
		)
		status.classList.add('status-error', 'status-visible')
		statusText.textContent = 'Error'

		responseElement.innerHTML = `<pre style="margin:0;font-family:inherit;">Error: ${error.message}</pre>`
		requestElement.innerHTML = ''
	}
}

async function handleUpdate() {
	const id = document.getElementById('update-id').value.trim()
	const title = document.getElementById('update-title').value.trim()
	const description = document.getElementById('update-description').value.trim()
	const priority = document.querySelector('.custom-select-task').value

	const responseElement = document.getElementById('response-text')
	const requestElement = document.getElementById('requests-text')
	const status = document.querySelector('.response-status')
	const statusText = document.querySelector('.response-status-text')

	try {
		const data = await Update(id, title, description, priority)

		status.classList.remove(
			'status-unvisible',
			'status-success',
			'status-error'
		)
		if (data.response && data.response.status === 'successful') {
			status.classList.add('status-success', 'status-visible')
			statusText.textContent = 'Success'
		} else {
			status.classList.add('status-error', 'status-visible')
			statusText.textContent = 'Error'
		}

		requestElement.innerHTML = `<pre style="margin:0;font-family:inherit;">${
			data.request.method
		} ${JSON.stringify(data.request.body, null, 2)}</pre>`
		responseElement.innerHTML = `<pre style="margin:0;font-family:inherit;">${JSON.stringify(
			data.response,
			null,
			2
		)}</pre>`
	} catch (error) {
		status.classList.remove(
			'status-unvisible',
			'status-success',
			'status-error'
		)
		status.classList.add('status-error', 'status-visible')
		statusText.textContent = 'Error'

		responseElement.innerHTML = `<pre style="margin:0;font-family:inherit;">Error: ${error.message}</pre>`
		requestElement.innerHTML = ''
	}
}

async function handleDelete() {
	const id = document.getElementById('delete-id').value.trim()
	const responseElement = document.getElementById('response-text')
	const requestElement = document.getElementById('requests-text')
	const status = document.querySelector('.response-status')
	const statusText = document.querySelector('.response-status-text')

	try {
		const data = await Delete(id)

		status.classList.remove(
			'status-unvisible',
			'status-success',
			'status-error'
		)
		if (data.response && data.response.status === 'successful') {
			status.classList.add('status-success', 'status-visible')
			statusText.textContent = 'Success'
		} else {
			status.classList.add('status-error', 'status-visible')
			statusText.textContent = 'Error'
		}

		requestElement.innerHTML = `<pre style="margin:0;font-family:inherit;">${data.request.method}</pre>`
		responseElement.innerHTML = `<pre style="margin:0;font-family:inherit;">${JSON.stringify(
			data.response,
			null,
			2
		)}</pre>`
	} catch (error) {
		status.classList.remove(
			'status-unvisible',
			'status-success',
			'status-error'
		)
		status.classList.add('status-error', 'status-visible')
		statusText.textContent = 'Error'

		responseElement.innerHTML = `<pre style="margin:0;font-family:inherit;">Error: ${error.message}</pre>`
		requestElement.innerHTML = ''
	}
}

async function handleClose() {
	const id = document.getElementById('close-id').value.trim()
	const responseElement = document.getElementById('response-text')
	const requestElement = document.getElementById('requests-text')
	const status = document.querySelector('.response-status')
	const statusText = document.querySelector('.response-status-text')

	try {
		const data = await Close(id)

		status.classList.remove(
			'status-unvisible',
			'status-success',
			'status-error'
		)
		if (data.response && data.response.status === 'successful') {
			status.classList.add('status-success', 'status-visible')
			statusText.textContent = 'Success'
		} else {
			status.classList.add('status-error', 'status-visible')
			statusText.textContent = 'Error'
		}

		requestElement.innerHTML = `<pre style="margin:0;font-family:inherit;">${data.request.method}</pre>`
		responseElement.innerHTML = `<pre style="margin:0;font-family:inherit;">${JSON.stringify(
			data.response,
			null,
			2
		)}</pre>`
	} catch (error) {
		status.classList.remove(
			'status-unvisible',
			'status-success',
			'status-error'
		)
		status.classList.add('status-error', 'status-visible')
		statusText.textContent = 'Error'

		responseElement.innerHTML = `<pre style="margin:0;font-family:inherit;">Error: ${error.message}</pre>`
		requestElement.innerHTML = ''
	}
}

async function handleSend() {
	const op = getActiveOperation()

	switch (op) {
		case 'btnradio1':
			await handleCreate()
			break
		case 'btnradio2':
			await handleGet()
			break
		case 'btnradio3':
			await handleUpdate()
			break
		case 'btnradio4':
			await handleDelete()
			break
		case 'btnradio5':
			await handleClose()
			break
		default:
			console.error('Unknown operation')
	}
	await loadTasks()
}
