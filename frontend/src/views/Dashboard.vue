<template>
  <div class="main-layout">
    <WelcomeHeader />
  </div>

  <div class="dashboard-container">
    <div class="dashboard-content">
      
      <!-- БЛОК ПРОФИЛЯ -->
      <section class="user-profile-section card-surface">
        <div class="avatar-block">
          <div class="avatar-image"></div>
          <div class="user-name-wrapper">
            <span class="user-label">User</span>
            <h1 class="username">SelfDev_Hero</h1>
          </div>
        </div>

        <div class="xp-level-block">
          <div class="xp-info">
            <span class="lvl-text">Level 14</span>
            <span class="xp-count">1500 / 3000 XP</span>
          </div>
          <div class="xp-bar-container">
            <div class="xp-bar-bg"></div>
            <div class="xp-bar-fill" style="width: 50%"></div>
          </div>
        </div>

        <div class="quick-stats">
          <div class="stat-box">
            <span class="stat-n">12</span>
            <span class="stat-t">Perfect Days</span>
          </div>
          <div class="stat-box streak">
            <span class="stat-n">142</span>
            <span class="stat-t">Days Streak</span>
          </div>
        </div>
      </section>

      <!-- НИЖНЯЯ ЧАСТЬ: ПРИВЫЧКИ И СТАТИСТИКА -->
      <div class="main-grid">
        
        <!-- ЛЕВАЯ КОЛОНКА: ПРИВЫЧКИ С НАВИГАЦИЕЙ И АНИМАЦИЕЙ -->
        <section class="habits-container">
          <!-- Категории привычек (Nav Bar) -->
          <nav class="categories-nav">
            <button 
              v-for="category in categories" 
              :key="category.value"
              class="nav-tab"
              :class="{ active: currentCategory === category.value }"
              @click="currentCategory = category.value"
            >
              {{ category.label }}
            </button>
          </nav>

          <!-- 🔥 Новое "окно видимости" с эффектом плавного затухания по краям -->
          <div class="habits-fade-viewport">
            <div class="habits-scroll-window">
              <TransitionGroup name="habit-fade" tag="div" class="habits-wrapper-layout">
                <article 
                  v-for="habit in filteredHabits" 
                  :key="habit.id" 
                  class="habit-card card-surface" 
                  :class="habit.color"
                >
                  <div class="habit-header">
                    <div class="habit-title-group">
                      <h2>{{ habit.name }}</h2>
                      <span class="habit-status">42/365 cleared</span>
                    </div>
                    <button class="btn btn-primary btn-sm">Done</button>
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
                          v-for="n in 365" 
                          :key="n" 
                          class="cube" 
                          :data-level="getRandomLevel()"
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
import { ref, computed } from 'vue'
import WelcomeHeader from '@/components/Header/WelcomeHeader.vue'

const currentCategory = ref('all')

const categories = [
  { label: 'All Habits', value: 'all' },
  { label: 'Languages', value: 'languages' },
  { label: 'Coding', value: 'coding' },
  { label: 'Health', value: 'health' }
]

const habitsList = [
  { id: 1, name: 'English', color: 'green', category: 'languages' },
  { id: 2, name: 'Japanese', color: 'blue', category: 'languages' },
  { id: 3, name: 'Python', color: 'purple', category: 'coding' },
  { id: 4, name: 'Gym Training', color: 'green', category: 'health' },
  { id: 5, name: 'Read Books', color: 'blue', category: 'health' }
]

const filteredHabits = computed(() => {
  if (currentCategory.value === 'all') return habitsList
  return habitsList.filter(habit => habit.category === currentCategory.value)
})

const getRandomLevel = () => Math.floor(Math.random() * 5)
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

.lvl-text { color: var(--text-primary); }
.xp-count { color: var(--text-secondary); }

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

.streak .stat-n { color: var(--accent-primary); }

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
  padding: 24px;
  position: relative;
  transition: border-color 0.25s ease, box-shadow 0.25s ease;
  width: 100%;
  box-sizing: border-box;
}

.habit-card:hover {
  border-color: var(--border-medium);
  box-shadow: 0 8px 30px rgba(0,0,0,0.15);
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
  padding: 16px;
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
  overflow-x: auto;
  overflow-y: hidden;
  padding-bottom: 4px;
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
  grid-template-rows: repeat(7, 12px);
  grid-auto-flow: column;
  grid-auto-columns: 12px;
  gap: 5px;
}

.cube {
  width: 12px;
  height: 12px;
  background: var(--border-default);
  border-radius: 3px;
}

/* Цвета хитмапа */
.green [data-level="1"] { background: #0e4429; }
.green [data-level="2"] { background: #006d32; }
.green [data-level="3"] { background: #26a641; }
.green [data-level="4"] { background: #39d353; }

.blue [data-level="1"] { background: rgba(59, 130, 246, 0.25); }
.blue [data-level="2"] { background: rgba(59, 130, 246, 0.5); }
.blue [data-level="3"] { background: var(--accent-primary); opacity: 0.8; }
.blue [data-level="4"] { background: var(--accent-primary); }

.purple [data-level="1"] { background: #3d1a78; }
.purple [data-level="2"] { background: #6e40c9; }
.purple [data-level="3"] { background: #9b72ff; }
.purple [data-level="4"] { background: #d2a8ff; }

.heatmap-legend {
  display: flex;
  justify-content: flex-end;
  align-items: center;
  gap: 10px;
  margin-top: 12px;
  font-size: 11px;
  color: var(--text-secondary);
}
.l-cubes { display: flex; gap: 4px; }

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