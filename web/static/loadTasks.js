import { getAll } from './api.js'

export async function loadTasks() {
	const taskListContainer = document.querySelector('.task-list')
	const taskCountElement = document.querySelector('.task-count-block p')
	if (!taskListContainer) return

	taskListContainer.innerHTML = ''

	try {
		const data = await getAll()
		if (data.response && data.response.status === 'successful') {
			const tasks = data.response.data.Lib
			const taskIds = Object.keys(tasks)
			taskCountElement.textContent = taskIds.length

			for (const id in tasks) {
				const task = tasks[id]

				const taskBlock = document.createElement('div')
				taskBlock.className = `task task-${task.id} mt-md`

				taskBlock.innerHTML = `
					<div class="task-block task-header">
						<div class="task-title">
							<div class="task-id"><p>${task.id}</p></div>
							<p>${task.title}</p>
						</div>
						<div class="priority priority-${task.priority.toLowerCase()}">
							<span class="priority-status">${task.priority}</span>
						</div>
					</div>
					<div class="task-block task-footer">
						<p>${task.status.charAt(0).toUpperCase() + task.status.slice(1)}</p>
					</div>
				`

				taskListContainer.appendChild(taskBlock)
			}
		} else {
			taskListContainer.innerHTML = '<p>No tasks found</p>'
		}
	} catch (error) {
		console.error('Error loading tasks:', error)
		taskListContainer.innerHTML = `<p>Error loading tasks: ${error.message}</p>`
	}
}
