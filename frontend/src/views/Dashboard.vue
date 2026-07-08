<template>
  <div class="main-layout">
    <WelcomeHeader />
  </div>

  <div class="dashboard-container">
    <div class="dashboard-content">
      <div v-if="loading" class="card-surface" style="padding: 16px 24px">Loading dashboard...</div>

      <div v-else-if="error" class="card-surface" style="padding: 16px 24px; color: #ff8a8a">
        {{ error }}
      </div>

      <!-- БЛОК ПРОФИЛЯ -->
      <section v-if="!loading && !error" class="user-profile-section card-surface">
        <div class="avatar-block">
          <img class="avatar-image" :src="defaultAvatar" alt="User avatar" />
          <div class="user-name-wrapper">
            <span class="user-label">User</span>
            <h1 class="username">{{ user?.username || 'SelfDev_Hero' }}</h1>
          </div>
        </div>

        <div class="xp-level-block">
          <div class="xp-info">
            <span class="lvl-text">Level {{ level }}</span>
            <span class="xp-count">{{ xpCurrent }} / {{ xpNext }} XP</span>
          </div>
          <div class="xp-bar-container">
            <div class="xp-bar-bg"></div>
            <div class="xp-bar-fill" :style="{ width: `${xpProgress}%` }"></div>
          </div>
        </div>

        <div class="quick-stats">
          <div class="stat-box">
            <span class="stat-n">{{ perfectDays }}</span>
            <span class="stat-t">Perfect Days</span>
          </div>
          <div class="stat-box streak">
            <span class="stat-n">{{ streakDays }}</span>
            <span class="stat-t">Days Streak</span>
          </div>
        </div>
      </section>

      <!-- НИЖНЯЯ ЧАСТЬ: ПРИВЫЧКИ И СТАТИСТИКА -->
      <div v-if="!loading && !error" class="main-grid">
        <!-- ЛЕВАЯ КОЛОНКА: ПРИВЫЧКИ С НАВИГАЦИЕЙ И АНИМАЦИЕЙ -->
        <section class="habits-container">
          <div class="habits-toolbar card-surface">
            <button
              class="btn btn-primary btn-sm"
              type="button"
              @click="showCreateHabitForm = !showCreateHabitForm"
            >
              Create Habit
            </button>
          </div>

          <div
            v-if="showCreateHabitForm"
            class="habit-modal-overlay"
            @click.self="showCreateHabitForm = false"
          >
            <form class="habit-create-form card-surface" @submit.prevent="createHabit">
              <h3>Create Habit</h3>
              <input
                v-model="newHabit.name"
                class="habit-input"
                type="text"
                placeholder="Habit name"
                required
              />
              <input
                v-model="newHabit.description"
                class="habit-input"
                type="text"
                placeholder="Description"
              />
              <label class="habit-checkbox">
                <input v-model="newHabit.isGood" type="checkbox" />
                <span>Good habit</span>
              </label>
              <div class="habit-create-actions">
                <button class="btn btn-primary btn-sm" type="submit">Save</button>
                <button class="btn btn-sm" type="button" @click="showCreateHabitForm = false">
                  Cancel
                </button>
              </div>
            </form>
          </div>

          <!-- 🔥 Новое "окно видимости" с эффектом плавного затухания по краям -->
          <div class="habits-fade-viewport">
            <div class="habits-scroll-window">
              <TransitionGroup name="habit-fade" tag="div" class="habits-wrapper-layout">
                <article
                  v-for="habit in habits"
                  :key="habit.id"
                  class="habit-card card-surface"
                  :class="habit.color"
                >
                  <div class="habit-header">
                    <div class="habit-title-group">
                      <h2>{{ habit.name }}</h2>
                      <span class="habit-status">{{ habit.confirmedCount }}/365 cleared</span>
                    </div>
                    <button class="btn btn-primary btn-sm" @click="toggleHabit(habit)">
                      {{ isHabitDoneToday(habit) ? 'Cancel' : 'Done' }}
                    </button>
                  </div>

                  <!-- СЕТКА КУБИКОВ (HEATMAP НА 365 ДНЕЙ) -->
                  <div class="heatmap-wrapper">
                    <div class="days-labels">
                      <span>Mon</span><span>Wed</span><span>Fri</span><span>Sun</span>
                    </div>
                    <!-- Горизонтальный скролл для кубиков, если они не влезают -->
                    <div class="cubes-scroll-container">
                      <div class="cubes-grid">
                        <div
                          v-for="day in habit.heatmap"
                          :key="day.key"
                          class="cube"
                          :data-level="day.level"
                          @click="day.level === 4 ? cancelHabit(habit.id) : confirmHabit(habit.id)"
                        ></div>
                      </div>
                    </div>
                  </div>

                  <div class="heatmap-legend">
                    <span>Less</span>
                    <div class="l-cubes">
                      <div class="cube" data-level="0"></div>
                      <div class="cube" data-level="1"></div>
                      <div class="cube" data-level="2"></div>
                      <div class="cube" data-level="3"></div>
                      <div class="cube" data-level="4"></div>
                    </div>
                    <span>More</span>
                  </div>
                </article>
              </TransitionGroup>
            </div>
          </div>
        </section>

        <!-- ПРАВАЯ КОЛОНКА: ФИКСИРОВАННАЯ СТАТИСТИКА -->
        <aside class="right-stats">
          <div class="stat-card card-surface">
            <h3>Activity Balance</h3>
            <div class="placeholder-chart"></div>
          </div>
          <div class="stat-card card-surface">
            <h3>Weekly Progress</h3>
            <div class="placeholder-chart bar"></div>
          </div>
        </aside>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import WelcomeHeader from '@/components/Header/WelcomeHeader.vue'
import defaultAvatar from '@/assets/default-avatar.jpg'

interface User {
  user_id: string
  role: string
  username: string
  email: string
}

interface HeatmapDay {
  key: string
  level: number
}

interface Habit {
  id: string
  name: string
  description: string
  isGood: boolean
  color: string
  category: string
  confirmedDates: string[]
  confirmedCount: number
  heatmap: HeatmapDay[]
}

const currentCategory = ref('all')
const user = ref<User | null>(null)
const habits = ref<Habit[]>([])
const loading = ref(true)
const error = ref('')
const showCreateHabitForm = ref(false)
const newHabit = ref({
  name: '',
  description: '',
  isGood: true,
})

const dayMs = 24 * 60 * 60 * 1000

function toIsoDate(value: string | Date) {
  const date = typeof value === 'string' ? new Date(value) : value
  return new Date(Date.UTC(date.getFullYear(), date.getMonth(), date.getDate()))
    .toISOString()
    .slice(0, 10)
}

function buildHeatmap(completedDates: string[]) {
  const completed = new Set(completedDates.map(toIsoDate))

  const days: HeatmapDay[] = []

  const today = new Date()
  today.setHours(0, 0, 0, 0)

  // находим понедельник текущей недели
  const end = new Date(today)
  const day = end.getDay()
  const mondayOffset = day === 0 ? -6 : 1 - day

  end.setDate(end.getDate() + mondayOffset + 6)

  // берем ровно 52 недели назад
  const start = new Date(end)
  start.setDate(start.getDate() - 52 * 7 + 1)

  const cursor = new Date(start)

  while (cursor <= end) {
    const key = toIsoDate(cursor)

    // будущие дни пустые НЕ добавляем
    if (cursor <= today) {
      days.push({
        key,
        level: completed.has(key) ? 4 : 0,
      })
    }

    cursor.setDate(cursor.getDate() + 1)
  }

  return days
}

function detectCategory(name: string, description: string, isGood: boolean) {
  const text = `${name} ${description}`.toLowerCase()
  if (/(english|japanese|spanish|language|learn|study)/.test(text)) return 'languages'
  if (/(code|coding|python|js|ts|program|dev)/.test(text)) return 'coding'
  if (/(gym|run|train|sport|health|sleep|water|meditat)/.test(text)) return 'health'
  return isGood ? 'health' : 'coding'
}

function habitColor(category: string) {
  if (category === 'coding') return 'purple'
  if (category === 'languages') return 'blue'
  return 'green'
}

function isHabitDoneToday(habit: Habit) {
  const today = toIsoDate(new Date())
  return habit.confirmedDates.map(toIsoDate).includes(today)
}

async function fetchJson<T>(url: string, options: RequestInit = {}) {
  const response = await fetch(url, {
    credentials: 'include',
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...(options.headers || {}),
    },
  })

  if (!response.ok) {
    throw new Error(await response.text())
  }

  return response.json() as Promise<T>
}

async function fetchUser() {
  const data = await fetchJson<User>('/api/auth/me')
  user.value = data
}

async function fetchHabitDates(habitId: string) {
  const data = await fetchJson<{
    dates?: Array<{ date?: string; Date?: string }>
    Dates?: Array<{ date?: string; Date?: string }>
  }>(`/api/habit/${encodeURIComponent(habitId)}/confirmed`)
  return (data.dates || data.Dates || [])
    .map((item) => item.date || item.Date || '')
    .filter(Boolean)
}

async function refreshHabit(habitId: string) {
  const dates = await fetchHabitDates(habitId)
  applyHabitDates(habitId, dates)
}

function applyHabitDates(habitId: string, dates: string[]) {
  habits.value = habits.value.map((habit) =>
    habit.id === habitId
      ? {
          ...habit,
          confirmedDates: dates,
          confirmedCount: dates.length,
          heatmap: buildHeatmap(dates),
        }
      : habit,
  )
}

async function fetchHabits(userId: string) {
  if (!userId) {
    throw new Error('Missing user id')
  }

  const data = await fetchJson<{
    habits?: Array<{
      habit_id?: string
      HabitId?: string
      name: string
      Name?: string
      description?: string
      Description?: string
      is_good?: boolean
      IsGood?: boolean
    }>
    Habits?: Array<{
      habit_id?: string
      HabitId?: string
      name: string
      Name?: string
      description?: string
      Description?: string
      is_good?: boolean
      IsGood?: boolean
    }>
  }>(`/api/habit/${encodeURIComponent(userId)}`)
  const nextHabits = await Promise.all(
    (data.habits || data.Habits || []).map(async (habit) => {
      const id = habit.habit_id || habit.HabitId || ''
      const name = habit.name || habit.Name || ''
      const description = habit.description || habit.Description || ''
      const isGood = habit.is_good ?? habit.IsGood ?? false
      const category = detectCategory(name, description, isGood)
      const confirmedDates = id ? await fetchHabitDates(id) : []

      return {
        id,
        name,
        description,
        isGood,
        color: habitColor(category),
        category,
        confirmedDates,
        confirmedCount: confirmedDates.length,
        heatmap: buildHeatmap(confirmedDates),
      }
    }),
  )

  habits.value = nextHabits
}

async function confirmHabit(habitId: string) {
  await fetchJson(`/api/habit/${encodeURIComponent(habitId)}/confirm`, { method: 'POST' })
  applyHabitDates(habitId, [
    ...new Set([
      ...(habits.value.find((habit) => habit.id === habitId)?.confirmedDates || []),
      toIsoDate(new Date()),
    ]),
  ])
  await refreshHabit(habitId)
}

async function cancelHabit(habitId: string) {
  await fetchJson(`/api/habit/${encodeURIComponent(habitId)}/cancel`, { method: 'POST' })
  applyHabitDates(
    habitId,
    (habits.value.find((habit) => habit.id === habitId)?.confirmedDates || []).filter(
      (date) => toIsoDate(date) !== toIsoDate(new Date()),
    ),
  )
  await refreshHabit(habitId)
}

async function toggleHabit(habit: Habit) {
  if (isHabitDoneToday(habit)) {
    await cancelHabit(habit.id)
    return
  }

  await confirmHabit(habit.id)
}

async function createHabit() {
  if (!user.value) return

  await fetchJson('/api/habit', {
    method: 'POST',
    body: JSON.stringify({
      user_id: user.value.user_id,
      name: newHabit.value.name,
      description: newHabit.value.description,
      is_good: newHabit.value.isGood,
    }),
  })

  showCreateHabitForm.value = false
  newHabit.value = { name: '', description: '', isGood: true }
  await fetchHabits(user.value.user_id)
}

const perfectDays = computed(() =>
  habits.value.reduce((sum, habit) => sum + habit.confirmedCount, 0),
)
const streakDays = computed(() => Math.max(...habits.value.map((habit) => habit.confirmedCount), 0))
const level = computed(() => Math.max(1, Math.floor(perfectDays.value / 10) + 1))
const xpCurrent = computed(() => perfectDays.value * 100)
const xpNext = computed(() => level.value * 1000)
const xpProgress = computed(() => Math.min(100, Math.round((xpCurrent.value / xpNext.value) * 100)))

onMounted(async () => {
  loading.value = true
  error.value = ''

  try {
    await fetchUser()
    await fetchHabits(user.value?.user_id || '')
  } catch (err) {
    error.value = err instanceof Error ? err.message : 'Failed to load dashboard'
  } finally {
    loading.value = false
  }
})
</script>

<style scoped>
/* =========================================
   LAYOUT & CONTAINERS
========================================= */
.dashboard-container {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  position: relative;
  overflow: hidden;
  padding: 100px 24px 40px;
  box-sizing: border-box;
}

.dashboard-container::before {
  content: '';
  position: absolute;
  width: 800px;
  height: 800px;
  background: radial-gradient(circle, rgba(149, 162, 223, 0.08), transparent 60%);
  filter: blur(60px);
  top: 30%;
  left: 50%;
  transform: translate(-50%, -50%);
  z-index: 0;
  pointer-events: none;
}

.dashboard-content {
  position: relative;
  z-index: 2;
  max-width: 1400px;
  width: 100%;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
  gap: 24px;
  height: calc(100vh - 150px);
}

.card-surface {
  background: var(--surface);
  border: 1px solid var(--border-subtle);
  backdrop-filter: blur(14px);
  -webkit-backdrop-filter: blur(14px);
  border-radius: 16px;
  box-shadow: var(--shadow-md);
  color: var(--text-primary);
}

/* =========================================
   USER PROFILE SECTION
========================================= */
.user-profile-section {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 24px 32px;
  flex-shrink: 0;
}

.avatar-block {
  display: flex;
  align-items: center;
  gap: 20px;
}

.avatar-image {
  width: 80px;
  height: 80px;
  background: var(--border-medium);
  border-radius: 14px;
  border: 1px solid var(--border-subtle);
}

.user-label {
  font-size: 13px;
  color: var(--text-secondary);
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

.username {
  font-size: 28px;
  font-weight: 800;
  margin: 4px 0 0 0;
  color: var(--accent-primary);
  text-shadow: 0 0 10px rgba(149, 162, 223, 0.1);
}

.xp-level-block {
  width: 360px;
}

.xp-info {
  display: flex;
  justify-content: space-between;
  margin-bottom: 8px;
  font-size: 14px;
  font-weight: 700;
}

.lvl-text {
  color: var(--text-primary);
}
.xp-count {
  color: var(--text-secondary);
}

.xp-bar-container {
  position: relative;
  height: 12px;
}

.xp-bar-bg {
  position: absolute;
  width: 100%;
  height: 100%;
  background: var(--border-default);
  border-radius: 6px;
}

.xp-bar-fill {
  position: absolute;
  height: 100%;
  background: linear-gradient(90deg, var(--accent-primary), var(--accent-dark));
  border-radius: 6px;
  box-shadow: 0 0 12px var(--border-glow);
}

.quick-stats {
  display: flex;
  gap: 32px;
}

.stat-box {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.stat-n {
  font-size: 32px;
  font-weight: 800;
  color: var(--text-primary);
}

.stat-t {
  font-size: 11px;
  color: var(--text-secondary);
  text-transform: uppercase;
  margin-top: 2px;
}

.streak .stat-n {
  color: var(--accent-primary);
}

/* =========================================
   MAIN GRID & HABITS NAVIGATION
========================================= */
.main-grid {
  display: grid;
  grid-template-columns: 1fr 380px;
  gap: 24px;
  flex: 1;
  overflow: hidden;
  min-height: 0;
}

.habits-container {
  overflow: hidden;
  display: flex;
  flex-direction: column;
  gap: 16px;
  min-height: 0;
}

.habits-toolbar {
  display: flex;
  justify-content: flex-end;
  padding: 12px 16px;
}

.habit-modal-overlay {
  position: fixed;
  inset: 0;
  z-index: 1000;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 24px;
  background: rgba(2, 6, 23, 0.62);
  backdrop-filter: blur(8px);
  -webkit-backdrop-filter: blur(8px);
}

.habit-create-form {
  width: min(100%, 420px);
  padding: 24px;
  display: flex;
  flex-direction: column;
  gap: 14px;
}

.habit-create-form h3 {
  margin: 0;
  font-size: 18px;
  font-weight: 700;
}

.habit-input {
  width: 100%;
  box-sizing: border-box;
  padding: 12px 14px;
  border-radius: 10px;
  border: 1px solid var(--border-default);
  background: var(--surface);
  color: var(--text-primary);
}

.habit-checkbox {
  display: flex;
  align-items: center;
  gap: 8px;
  color: var(--text-secondary);
  font-size: 14px;
}

.habit-create-actions {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}

.categories-nav {
  display: flex;
  gap: 8px;
  padding: 4px;
  background: var(--surface);
  border: 1px solid var(--border-subtle);
  border-radius: 12px;
  align-self: flex-start;
  flex-shrink: 0;
}

.nav-tab {
  background: transparent;
  border: none;
  color: var(--text-secondary);
  padding: 8px 16px;
  font-size: 14px;
  font-weight: 700;
  font-family: 'Evolventa', sans-serif;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.2s ease;
}

.nav-tab:hover {
  color: var(--text-primary);
}

.nav-tab.active {
  background: var(--accent-primary);
  color: var(--bg-primary);
}

/* =========================================
   🔥 ДОПОЛНИТЕЛЬНОЕ ОКНО ВИДИМОСТИ (FADE VIEWPORT)
========================================= */
.habits-fade-viewport {
  flex: 1;
  min-height: 0;
  position: relative;
  overflow: hidden;
  /* Мягкая CSS-маска. Края плавно растворяются в прозрачность на 24px сверху и снизу */
  -webkit-mask-image: linear-gradient(
    to bottom,
    transparent 0%,
    black 24px,
    black calc(100% - 24px),
    transparent 100%
  );
  mask-image: linear-gradient(
    to bottom,
    transparent 0%,
    black 24px,
    black calc(100% - 24px),
    transparent 100%
  );
}

.habits-scroll-window {
  height: 100%;
  overflow-y: auto;
  overflow-x: hidden;
  padding: 24px 8px 24px 0; /* Внутренние отступы компенсируют зону маски */
  box-sizing: border-box;
}

.habits-wrapper-layout {
  display: flex;
  flex-direction: column;
  gap: 20px;
  position: relative;
}

/* Стилизация скроллбара */
.habits-scroll-window::-webkit-scrollbar {
  width: 6px;
}
.habits-scroll-window::-webkit-scrollbar-thumb {
  background: var(--border-medium);
  border-radius: 10px;
}
.habits-scroll-window::-webkit-scrollbar-thumb:hover {
  background: var(--border-strong);
}

/* =========================================
   АНИМАЦИЯ ВЫЛЕТА КАРТОЧЕК (Vue Transitions)
========================================= */
.habit-fade-enter-active,
.habit-fade-leave-active {
  transition: all 0.35s cubic-bezier(0.4, 0, 0.2, 1);
}

.habit-fade-enter-from {
  opacity: 0;
  transform: translateY(24px) scale(0.98);
}

.habit-fade-leave-to {
  opacity: 0;
  transform: translateY(-24px) scale(0.96);
}

.habit-fade-leave-active {
  position: absolute;
  left: 0;
  right: 0;
  z-index: 0;
  pointer-events: none;
}

.habit-fade-move {
  transition: transform 0.35s cubic-bezier(0.4, 0, 0.2, 1);
}

/* =========================================
   HABIT CARD & HEATMAP (365 ДНЕЙ)
========================================= */
.habit-card {
  padding: 20px;
  position: relative;
  transition:
    border-color 0.25s ease,
    box-shadow 0.25s ease;
  width: 100%;
  box-sizing: border-box;
}

.habit-card:hover {
  border-color: var(--border-medium);
  box-shadow: 0 8px 30px rgba(0, 0, 0, 0.15);
}

.habit-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.habit-header h2 {
  font-size: 24px;
  font-weight: 700;
  margin: 0;
}

.habit-status {
  font-size: 13px;
  color: var(--text-secondary);
  display: block;
  margin-top: 4px;
}

.heatmap-wrapper {
  display: flex;
  gap: 16px;
  background: rgba(0, 0, 0, 0.15);
  padding: 12px;
  border-radius: 12px;
  border: 1px solid var(--border-subtle);
  width: 100%;
  box-sizing: border-box;
}

.days-labels {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  font-size: 10px;
  color: var(--text-secondary);
  padding: 2px 0;
  flex-shrink: 0;
}

/* Контейнер для горизонтального скролла сетки кубиков (365 дней) */
.cubes-scroll-container {
  flex: 1;
  overflow: hidden;
  width: 100%;
}

.cubes-grid::after {
  content: '';
}

/* Стилизация скроллбара для хитмапа */
.cubes-scroll-container::-webkit-scrollbar {
  height: 4px;
}
.cubes-scroll-container::-webkit-scrollbar-thumb {
  background: var(--border-subtle);
  border-radius: 4px;
}

/* 🔥 ИСПРАВЛЕНИЕ: Кубики сохраняют исходный размер (12px), сетка вмещает весь год */
.cubes-grid {
  display: grid;
  grid-template-rows: repeat(7, minmax(0, 1fr));
  grid-auto-flow: column;
  grid-auto-columns: minmax(0, 1fr);
  gap: 2px;
  width: 100%;
}

.cube {
  width: 100%;
  aspect-ratio: 1;
  background: var(--border-default);
  border-radius: 3px;
}

/* Цвета хитмапа */
.green [data-level='1'] {
  background: #0e4429;
}
.green [data-level='2'] {
  background: #006d32;
}
.green [data-level='3'] {
  background: #26a641;
}
.green [data-level='4'] {
  background: #39d353;
}

.blue [data-level='1'] {
  background: rgba(59, 130, 246, 0.25);
}
.blue [data-level='2'] {
  background: rgba(59, 130, 246, 0.5);
}
.blue [data-level='3'] {
  background: var(--accent-primary);
  opacity: 0.8;
}
.blue [data-level='4'] {
  background: var(--accent-primary);
}

.purple [data-level='1'] {
  background: #3d1a78;
}
.purple [data-level='2'] {
  background: #6e40c9;
}
.purple [data-level='3'] {
  background: #9b72ff;
}
.purple [data-level='4'] {
  background: #d2a8ff;
}

.heatmap-legend {
  display: flex;
  justify-content: flex-end;
  align-items: center;
  gap: 10px;
  margin-top: 12px;
  font-size: 11px;
  color: var(--text-secondary);
}
.l-cubes {
  display: flex;
  gap: 4px;
}

/* =========================================
   RIGHT SIDEBAR & CHARTS
========================================= */
.right-stats {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.stat-card {
  padding: 24px;
  flex: 1;
  display: flex;
  flex-direction: column;
}

.stat-card h3 {
  font-size: 18px;
  font-weight: 700;
  margin: 0 0 16px 0;
}

.placeholder-chart {
  flex: 1;
  min-height: 120px;
  background: rgba(0, 0, 0, 0.1);
  border-radius: 12px;
  border: 1px dashed var(--border-medium);
}

/* =========================================
   BUTTONS
========================================= */
.btn {
  padding: 10px 20px;
  border-radius: 10px;
  font-weight: 700;
  font-family: 'Evolventa', sans-serif;
  cursor: pointer;
  transition: 0.25s ease;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  border: none;
}

.btn-sm {
  padding: 8px 18px;
  font-size: 14px;
}

.btn-primary {
  background: linear-gradient(135deg, var(--accent-primary), var(--accent-dark));
  color: white;
  box-shadow: 0 4px 15px rgba(59, 130, 246, 0.2);
}

.btn-primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(59, 130, 246, 0.3);
}

/* =========================================
   RESPONSIVE
========================================= */
@media (max-width: 1024px) {
  .dashboard-content {
    height: auto;
    overflow: visible;
  }
  .main-grid {
    grid-template-columns: 1fr;
  }
  .habits-scroll-window {
    overflow-y: visible;
    height: auto;
  }
  .user-profile-section {
    flex-direction: column;
    gap: 20px;
    align-items: flex-start;
  }
  .xp-level-block {
    width: 100%;
  }
}
</style>
