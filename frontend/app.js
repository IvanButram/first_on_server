const API_BASE_URL = "";

const taskForm = document.getElementById("task-form");
const titleInput = document.getElementById("title");
const descriptionInput = document.getElementById("description");
const taskList = document.getElementById("task-list");
const statusText = document.getElementById("status");
const reloadBtn = document.getElementById("reload-btn");

function setStatus(message, isError = false) {
  statusText.textContent = message;
  statusText.style.color = isError ? "#d73a49" : "#555";
}

async function parseJsonSafe(response) {
  const text = await response.text();
  if (!text) return null;
  return JSON.parse(text);
}

function formatDate(dateString) {
  if (!dateString) return "—";
  return new Date(dateString).toLocaleString("ru-RU");
}

async function getTasks() {
  const response = await fetch(`${API_BASE_URL}/tasks`);

  if (!response.ok) {
    const errorData = await parseJsonSafe(response);
    throw new Error(errorData?.Message || "Не удалось загрузить задачи");
  }

  const data = await response.json();
  return data.Tasks || [];
}

async function createTask(title, description) {
  const response = await fetch(`${API_BASE_URL}/tasks`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      title,
      description,
    }),
  });

  if (!response.ok) {
    const errorData = await parseJsonSafe(response);
    throw new Error(errorData?.Message || "Не удалось создать задачу");
  }

  return response.json();
}

async function completeTask(title) {
  const response = await fetch(
    `${API_BASE_URL}/tasks/${encodeURIComponent(title)}`,
    {
      method: "PATCH",
    }
  );

  if (!response.ok) {
    const errorData = await parseJsonSafe(response);
    throw new Error(errorData?.Message || "Не удалось отметить задачу выполненной");
  }

  return response.json();
}

async function deleteTask(title) {
  const response = await fetch(
    `${API_BASE_URL}/tasks/${encodeURIComponent(title)}`,
    {
      method: "DELETE",
    }
  );

  if (!response.ok) {
    const errorData = await parseJsonSafe(response);
    throw new Error(errorData?.Message || "Не удалось удалить задачу");
  }

  if (response.status === 204) {
    return;
  }

  return parseJsonSafe(response);
}

function createTaskElement(task) {
  const li = document.createElement("li");
  li.className = `task-item ${task.Completed ? "done" : ""}`;

  const description = task.Description || "Без описания";

  li.innerHTML = `
    <div class="task-header">
      <div>
        <h3 class="task-title">${task.Title}</h3>
        <p class="task-meta">ID: ${task.Id}</p>
        <p class="task-meta">Создано: ${formatDate(task.CreatedAt)}</p>
        <p class="task-meta">Завершено: ${formatDate(task.CompletedAt)}</p>
      </div>
      <div class="task-actions">
        ${task.Completed ? "" : `<button class="btn-secondary complete-btn">Готово</button>`}
        <button class="btn-danger delete-btn">Удалить</button>
      </div>
    </div>
    <p class="task-description">${description}</p>
  `;

  const completeBtn = li.querySelector(".complete-btn");
  const deleteBtn = li.querySelector(".delete-btn");

  if (completeBtn) {
    completeBtn.addEventListener("click", async () => {
      try {
        setStatus("Отмечаю задачу выполненной...");
        await completeTask(task.Title);
        await loadTasks();
        setStatus(`Задача "${task.Title}" завершена`);
      } catch (error) {
        setStatus(error.message, true);
      }
    });
  }

  deleteBtn.addEventListener("click", async () => {
    try {
      setStatus("Удаляю задачу...");
      await deleteTask(task.Title);
      await loadTasks();
      setStatus(`Задача "${task.Title}" удалена`);
    } catch (error) {
      setStatus(error.message, true);
    }
  });

  return li;
}

function renderTasks(tasks) {
  taskList.innerHTML = "";

  if (!tasks.length) {
    const empty = document.createElement("li");
    empty.className = "empty";
    empty.textContent = "Пока задач нет";
    taskList.appendChild(empty);
    return;
  }

  tasks.forEach((task) => {
    taskList.appendChild(createTaskElement(task));
  });
}

async function loadTasks() {
  try {
    setStatus("Загружаю задачи...");
    const tasks = await getTasks();
    renderTasks(tasks);
    setStatus(`Загружено задач: ${tasks.length}`);
  } catch (error) {
    taskList.innerHTML = "";
    setStatus(error.message, true);
  }
}

taskForm.addEventListener("submit", async (event) => {
  event.preventDefault();

  const title = titleInput.value.trim();
  const description = descriptionInput.value.trim();

  if (!title) {
    setStatus("Название задачи обязательно", true);
    return;
  }

  try {
    setStatus("Создаю задачу...");
    await createTask(title, description);
    taskForm.reset();
    await loadTasks();
    setStatus(`Задача "${title}" создана`);
  } catch (error) {
    setStatus(error.message, true);
  }
});

reloadBtn.addEventListener("click", async () => {
  await loadTasks();
});

document.addEventListener("DOMContentLoaded", async () => {
  await loadTasks();
});