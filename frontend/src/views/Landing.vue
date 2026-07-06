<template>
  <div class="main-layout">
    <LandingHeader />
  </div>

  <main class="landing-page">
    <section class="hero">
      <div class="hero-content">
        <p class="eyebrow">SelfDev</p>
        <h1>Track everything that matters.</h1>
        <p class="hero-copy">
          A clean self-development tracker for habits, tasks, and goals. Create your account,
          verify your email, and keep your routine in one focused place.
        </p>

        <div class="hero-actions">
          <RouterLink to="/register" class="btn btn-primary">Create account</RouterLink>
        </div>
      </div>

      <div class="hero-panel">
        <div class="stat-card">
          <span class="stat-label">Flow</span>
          <strong>Register, verify, return</strong>
          <p>Simple onboarding built around email verification and secure sign-in.</p>
        </div>

        <div class="stat-card stat-card-accent">
          <span class="stat-label">Focus</span>
          <strong>Less clutter, more consistency</strong>
          <p>A calm interface with theme support, soft glass cards, and quick feedback.</p>
        </div>
      </div>

      <div class="hero-decoration"></div>
    </section>

    <section id="about" class="content-section">
      <div class="section-head">
        <p class="eyebrow">About</p>
        <h2>Built for daily progress without the noise.</h2>
      </div>

      <div class="section-card section-card-large">
        <p>
          SelfDev is a personal tracker designed to stay out of your way. The current product
          centers on account creation, email verification, secure login, and a private profile
          area so you can return to the same workspace every day.
        </p>
        <p>
          The tone is intentionally minimal: clear actions, calm visuals, and a lightweight flow
          that supports repeat use rather than constant setup.
        </p>
      </div>
    </section>

    <section id="features" class="content-section">
      <div class="section-head">
        <p class="eyebrow">Features</p>
        <h2>Everything the current build actually does.</h2>
      </div>

      <div class="feature-grid">
        <article v-for="feature in features" :key="feature.title" class="feature-card">
          <h3>{{ feature.title }}</h3>
          <p>{{ feature.text }}</p>
        </article>
      </div>
    </section>

    <section id="tech" class="content-section">
      <div class="section-head">
        <p class="eyebrow">Tech Stack</p>
        <h2>Modern tools, kept deliberately small.</h2>
      </div>

      <div class="tech-grid">
        <article v-for="group in techStack" :key="group.title" class="section-card tech-card">
          <h3>{{ group.title }}</h3>
          <p>{{ group.text }}</p>
        </article>
      </div>
    </section>

    <section id="faq" class="content-section">
      <div class="section-head">
        <p class="eyebrow">FAQ</p>
        <h2>Common questions, answered plainly.</h2>
      </div>

      <div class="faq-grid">
        <article v-for="item in faqs" :key="item.question" class="faq-card">
          <h3>{{ item.question }}</h3>
          <p>{{ item.answer }}</p>
        </article>
      </div>
    </section>

    <section class="cta-section">
      <div class="cta-card">
        <p class="eyebrow">Final CTA</p>
        <h2>Start your routine with less friction.</h2>
        <p>
          Create your account, verify your email, and keep your habits, tasks, and goals in one
          calm place.
        </p>

        <div class="hero-actions">
          <RouterLink to="/register" class="btn btn-primary">Create account</RouterLink>
        </div>
      </div>
    </section>
  </main>
</template>

<script setup lang="ts">
import LandingHeader from '@/components/Header/LandingHeader.vue'

const features = [
  {
    title: 'Habit, task, and goal tracking',
    text: 'The product is positioned as a personal tracker for everyday self-development.',
  },
  {
    title: 'Username or email login',
    text: 'Users can sign in with either identifier, which keeps the login form simple.',
  },
  {
    title: 'Email verification',
    text: 'Registration triggers a verification flow before login is allowed.',
  },
  {
    title: 'Secure session cookies',
    text: 'The backend sets HttpOnly access and refresh cookies after a successful login.',
  },
  {
    title: 'Token refresh support',
    text: 'Refresh tokens are stored server-side and used to issue new access tokens.',
  },
  {
    title: 'Persistent light and dark mode',
    text: 'Theme choice is saved in localStorage and applied across the app.',
  },
  {
    title: 'Responsive glassmorphism UI',
    text: 'The interface uses soft surfaces, blur, and responsive spacing across devices.',
  },
  {
    title: 'Toast-based feedback',
    text: 'Registration, login, and verification states surface through brief notifications.',
  },
  {
    title: 'Container-ready deployment',
    text: 'The project includes Docker, compose files, and Nginx configs for delivery.',
  },
]

const techStack = [
  {
    title: 'Frontend',
    text: 'Vue 3, Vue Router, TypeScript, Vite, and Vue Toastification.',
  },
  {
    title: 'Backend',
    text: 'Go 1.26.2 with JWT, pgx, Redis, bcrypt, Zap, and the Resend email API.',
  },
  {
    title: 'Data and delivery',
    text: 'PostgreSQL 18, Redis, Nginx, Docker, docker-compose, and Task.',
  },
]

const faqs = [
  {
    question: 'What is SelfDev?',
    answer: 'A personal tracker built around habits, tasks, and goals with a clean auth flow.',
  },
  {
    question: 'Do I need an account?',
    answer: 'Yes. The app is centered on registration, verification, and sign-in.',
  },
  {
    question: 'Can I log in with either username or email?',
    answer: 'Yes. The login form accepts both formats.',
  },
  {
    question: 'Why do I need to verify my email?',
    answer: 'The backend requires verification before login is allowed.',
  },
  {
    question: 'What happens after registration?',
    answer: 'You are sent back to the login screen and prompted to verify your email.',
  },
  {
    question: 'How are sessions handled?',
    answer: 'The backend sets secure HttpOnly cookies for access and refresh tokens.',
  },
  {
    question: 'Does the app remember my theme preference?',
    answer: 'Yes. The current theme is stored in localStorage.',
  },
  {
    question: 'Is the UI responsive?',
    answer: 'Yes. The layout and header adapt to smaller screens.',
  },
  {
    question: 'What powers the backend?',
    answer: 'Go, PostgreSQL, Redis, JWT, bcrypt, Zap, and Resend.',
  },
  {
    question: 'Can I self-host it?',
    answer: 'Yes. The repository includes Docker and docker-compose setup.',
  },
  {
    question: 'Is this a single-page app?',
    answer: 'Yes. The frontend uses Vue Router for navigation.',
  },
  {
    question: 'What kind of design does it use?',
    answer: 'A minimal glassmorphism style with soft gradients and strong typography.',
  },
]
</script>

<style scoped>
/* =========================
   PAGE SHELL
========================= */
.landing-page {
  width: 100%;
  padding-bottom: 20px;
}

/* =========================
   SHARED SECTION STYLES
========================= */
.content-section,
.cta-section,
.hero {
  position: relative;
  width: min(1120px, calc(100% - 32px));
  margin: 0 auto;
}

.content-section {
  padding: 28px 0 0;
  scroll-margin-top: 96px;
}

.section-head {
  display: flex;
  flex-direction: column;
  gap: 10px;
  margin-bottom: 18px;
}

.eyebrow {
  font-size: 12px;
  font-weight: 700;
  letter-spacing: 0.14em;
  text-transform: uppercase;
  color: var(--accent-primary);
}

.section-head h2,
.cta-card h2 {
  font-size: clamp(28px, 3.4vw, 44px);
  line-height: 1.1;
  color: var(--text-primary);
}

.section-card,
.feature-card,
.faq-card,
.cta-card,
.hero-content,
.hero-panel {
  background: var(--surface);
  border: 1px solid var(--surface-border);
  backdrop-filter: blur(14px);
  -webkit-backdrop-filter: blur(14px);
  box-shadow: var(--shadow-md);
  border-radius: 16px;
}

/* =========================
   HERO LAYOUT
========================= */
.hero {
  min-height: 100vh;
  display: grid;
  grid-template-columns: minmax(0, 1.2fr) minmax(280px, 0.8fr);
  gap: 24px;
  align-items: center;
  padding: 92px 0 12px;
  overflow: hidden;
}

/* =========================
   CONTENT CARD
========================= */
.hero-content {
  display: flex;
  flex-direction: column;
  gap: 16px;
  padding: clamp(24px, 3vw, 40px);
  z-index: 2;
}

.hero-panel {
  display: grid;
  gap: 16px;
  padding: clamp(18px, 2vw, 24px);
  z-index: 2;
}

.stat-card {
  padding: 18px;
  border-radius: 14px;
  border: 1px solid var(--border-default);
  background: rgba(255, 255, 255, 0.03);
}

.stat-card-accent {
  background: linear-gradient(135deg, rgba(100, 200, 255, 0.14), rgba(108, 123, 255, 0.12));
}

.stat-label {
  display: inline-block;
  margin-bottom: 10px;
  font-size: 12px;
  font-weight: 700;
  letter-spacing: 0.12em;
  text-transform: uppercase;
  color: var(--accent-primary);
}

.stat-card strong,
.feature-card h3,
.tech-card h3,
.faq-card h3 {
  display: block;
  margin-bottom: 8px;
  font-size: 18px;
  color: var(--text-primary);
}

.stat-card p,
.section-card p,
.feature-card p,
.tech-card p,
.faq-card p,
.cta-card p,
.hero-copy {
  color: var(--text-secondary);
  line-height: 1.6;
}

.section-card-large {
  padding: clamp(20px, 3vw, 30px);
  display: grid;
  gap: 12px;
}

/* =========================
   TITLE (RESPONSIVE FONTS)
========================= */
.hero h1 {
  font-size: clamp(36px, 6vw, 64px);
  font-weight: 800;
  line-height: 1.1;
  color: var(--text-primary);
  text-shadow:
    0 0 10px rgba(149, 162, 223, 0.18),
    0 0 25px rgba(59, 130, 246, 0.12);
}

.hero-copy {
  max-width: 560px;
  font-size: clamp(15px, 1.4vw, 18px);
}

/* =========================
   ACTION BUTTONS
========================= */
.hero-actions {
  display: flex;
  gap: 14px;
  flex-wrap: wrap;
  margin-top: 10px;
}

.btn {
  padding: 12px 24px;
  border-radius: 10px;
  font-weight: 700;
  text-decoration: none;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  transition: all 0.25s ease;
  font-family: 'Hind Madurai', sans-serif;
}

.btn-primary {
  background: linear-gradient(135deg, var(--accent-primary), var(--accent-dark));
  color: white;
  box-shadow:
    0 10px 30px rgba(59, 130, 246, 0.25),
    0 0 20px rgba(149, 162, 223, 0.15);
}

.btn-primary:hover {
  transform: translateY(-3px);
  box-shadow:
    0 15px 40px rgba(59, 130, 246, 0.35),
    0 0 30px rgba(149, 162, 223, 0.25);
}

.btn-secondary {
  background: var(--surface);
  color: var(--text-primary);
  border: 1px solid var(--surface-border);
  backdrop-filter: blur(10px);
}

.btn-secondary:hover {
  transform: translateY(-3px);
  border-color: var(--accent-primary);
  box-shadow: 0 0 25px rgba(149, 162, 223, 0.15);
}

/* =========================
   GRID SECTIONS
========================= */
.feature-grid,
.faq-grid,
.tech-grid {
  display: grid;
  gap: 16px;
}

.feature-grid,
.tech-grid {
  grid-template-columns: repeat(3, minmax(0, 1fr));
}

.faq-grid {
  grid-template-columns: repeat(2, minmax(0, 1fr));
}

.feature-card,
.faq-card,
.tech-card {
  padding: 20px;
}

/* =========================
   CTA
========================= */
.cta-section {
  padding: 44px 0 0;
}

.cta-card {
  padding: clamp(24px, 3vw, 36px);
  display: flex;
  flex-direction: column;
  gap: 14px;
  align-items: flex-start;
}

/* =========================
   BACKGROUND GLOW (FIXED LAYERS)
========================= */
.hero::before {
  content: '';
  position: absolute;
  width: 600px;
  height: 600px;
  background: radial-gradient(circle, rgba(149, 162, 223, 0.12), transparent 60%);
  filter: blur(50px);
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  z-index: 0;
  pointer-events: none;
}

/* =========================
   RESPONSIVE
========================= */
@media (max-width: 1024px) {
  .hero {
    grid-template-columns: 1fr;
    padding-top: 84px;
  }

  .feature-grid,
  .tech-grid {
    grid-template-columns: repeat(2, minmax(0, 1fr));
  }
}

@media (max-width: 768px) {
  .content-section,
  .cta-section,
  .hero {
    width: min(1120px, calc(100% - 24px));
  }

  .hero {
    min-height: auto;
    padding: 96px 0 18px;
  }

  .hero-actions {
    flex-direction: column;
    width: 100%;
  }

  .btn {
    width: 100%;
  }

  .feature-grid,
  .tech-grid,
  .faq-grid {
    grid-template-columns: 1fr;
  }
}
</style>