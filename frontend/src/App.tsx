import { useState } from 'react'
import CreateTournamentModal from './components/CreateTournamentModal'

type Tab = 'tournaments' | 'leaderboard'

const API_BASE = import.meta.env.VITE_API_URL || '/api/v1'

async function createTournament(data: {
  name: string
  type: string
  status: string
  location: string
  season: string
  courts: string[]
  maxPoints: number
  players: { name: string; gender: string }[]
}): Promise<{ success: true; data: { id: string; name: string; type: string; status: string; location: string } } | { success: false; error: string }> {
  const res = await fetch(`${API_BASE}/tournaments`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({
      name: data.name,
      type: data.type,
      status: data.status,
      location: data.location,
      season: data.season,
      courts: data.courts.length ? data.courts : ['Court 1'],
      max_points: data.maxPoints || 15,
      players: data.players.map((p) => ({ name: p.name, gender: p.gender })),
    }),
  })
  let json: { success?: boolean; data?: { id: string; name: string; type: string; status: string; location: string }; error?: { message?: string } }
  try {
    json = await res.json()
  } catch {
    return { success: false, error: res.ok ? 'Invalid response' : res.statusText || 'Request failed' }
  }
  if (res.ok && json.success && json.data) {
    return { success: true, data: json.data }
  }
  return { success: false, error: json.error?.message || (res.ok ? 'Failed to create tournament' : res.statusText || 'Request failed') }
}

function App() {
  const [activeTab, setActiveTab] = useState<Tab>('tournaments')
  const [showCreateModal, setShowCreateModal] = useState(false)
  const [tournaments, setTournaments] = useState<Array<{ id: string; name: string; type: string; status: string; location: string }>>([])
  const [createError, setCreateError] = useState<string | null>(null)

  const handleCreateTournament = async (data: Parameters<typeof createTournament>[0]) => {
    setCreateError(null)
    try {
      const result = await createTournament(data)
      if (result.success) {
        setTournaments((prev) => [
          ...prev,
          {
            id: result.data.id,
            name: result.data.name,
            type: result.data.type,
            status: result.data.status,
            location: result.data.location,
          },
        ])
        setShowCreateModal(false)
      } else {
        setCreateError(result.error)
      }
    } catch (err) {
      setCreateError(err instanceof Error ? err.message : 'Connection failed')
    }
  }

  return (
    <div className="min-h-screen bg-gradient-to-b from-base-200 via-base-100 to-base-200">
      {/* Hero / Header */}
      <header className="relative overflow-hidden border-b border-base-300/50 bg-base-100 shadow-sm">
        <div className="absolute inset-0 bg-[radial-gradient(ellipse_80%_50%_at_50%_-20%,rgba(16,185,129,0.15),transparent)]" />
        <div className="container mx-auto relative px-4 py-8 md:py-12">
          <div className="flex flex-col items-center text-center gap-4">
            <h1 className="text-5xl md:text-6xl font-black tracking-tight text-emerald-600 italic drop-shadow-sm">
              BERPADEL
            </h1>
            <p className="text-base-content/80 text-sm md:text-base font-medium max-w-md">
              Berlin padel community — tournaments, leaderboards & fun games
            </p>
            <button
              onClick={() => setShowCreateModal(true)}
              className="btn btn-primary btn-lg gap-2 shadow-lg shadow-primary/25 hover:shadow-primary/40 transition-all"
            >
              <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                <path fillRule="evenodd" d="M10 3a1 1 0 011 1v5h5a1 1 0 110 2h-5v5a1 1 0 11-2 0v-5H4a1 1 0 110-2h5V4a1 1 0 011-1z" clipRule="evenodd" />
              </svg>
              Create Tournament
            </button>
          </div>
        </div>
      </header>

      {/* Tabs + Content */}
      <main className="container mx-auto px-4 py-6 md:py-10">
        <div role="tablist" className="tabs tabs-boxed bg-base-200/80 p-1 rounded-xl w-fit mx-auto mb-8">
          <button
            role="tab"
            className={`tab ${activeTab === 'tournaments' ? 'tab-active' : ''}`}
            onClick={() => setActiveTab('tournaments')}
          >
            Tournaments
          </button>
          <button
            role="tab"
            className={`tab ${activeTab === 'leaderboard' ? 'tab-active' : ''}`}
            onClick={() => setActiveTab('leaderboard')}
          >
            Leaderboard
          </button>
        </div>

        {activeTab === 'tournaments' && (
          <section className="max-w-2xl mx-auto">
            {tournaments.length === 0 ? (
              <div className="card bg-base-100 shadow-xl border border-base-300/50">
                <div className="card-body items-center text-center py-16">
                  <div className="w-20 h-20 rounded-full bg-emerald-500/10 flex items-center justify-center mb-4">
                    <svg xmlns="http://www.w3.org/2000/svg" className="h-10 w-10 text-emerald-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" />
                    </svg>
                  </div>
                  <h2 className="card-title text-xl">No tournaments yet</h2>
                  <p className="text-base-content/70 text-sm">Start your first tournament and get the community playing.</p>
                  <button
                    onClick={() => setShowCreateModal(true)}
                    className="btn btn-primary mt-4"
                  >
                    Create Tournament
                  </button>
                </div>
              </div>
            ) : (
              <div className="space-y-4">
                {tournaments.map((t) => (
                  <div key={t.id} className="card bg-base-100 shadow-md border border-base-300/50 hover:shadow-lg transition-shadow">
                    <div className="card-body">
                      <div className="flex flex-wrap items-center justify-between gap-2">
                        <h3 className="font-bold text-lg">{t.name}</h3>
                        <span className={`badge badge-sm ${t.status === 'ongoing' ? 'badge-success' : t.status === 'completed' ? 'badge-ghost' : 'badge-warning'}`}>
                          {t.status}
                        </span>
                      </div>
                      <p className="text-sm text-base-content/70">{t.location} · {t.type.replace(/_/g, ' ')}</p>
                    </div>
                  </div>
                ))}
              </div>
            )}
          </section>
        )}

        {activeTab === 'leaderboard' && (
          <section className="max-w-2xl mx-auto">
            <div className="card bg-base-100 shadow-xl border border-base-300/50">
              <div className="card-body items-center text-center py-16">
                <div className="w-20 h-20 rounded-full bg-amber-500/10 flex items-center justify-center mb-4">
                  <svg xmlns="http://www.w3.org/2000/svg" className="h-10 w-10 text-amber-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 12l2 2 4-4M7.835 4.697a3.42 3.42 0 001.946-.806 3.42 3.42 0 014.438 0 3.42 3.42 0 001.946.806 3.42 3.42 0 013.138 3.138 3.42 3.42 0 00.806 1.946 3.42 3.42 0 010 4.438 3.42 3.42 0 00-.806 1.946 3.42 3.42 0 01-3.138 3.138 3.42 3.42 0 00-1.946.806 3.42 3.42 0 01-4.438 0 3.42 3.42 0 00-1.946-.806 3.42 3.42 0 01-3.138-3.138 3.42 3.42 0 00-.806-1.946 3.42 3.42 0 010-4.438 3.42 3.42 0 00.806-1.946 3.42 3.42 0 013.138-3.138z" />
                  </svg>
                </div>
                <h2 className="card-title text-xl">Leaderboard</h2>
                <p className="text-base-content/70 text-sm">Community rankings will appear here after tournaments are played.</p>
              </div>
            </div>
          </section>
        )}
      </main>

      <CreateTournamentModal
        open={showCreateModal}
        onClose={() => { setShowCreateModal(false); setCreateError(null) }}
        onSubmit={handleCreateTournament}
        error={createError}
      />
    </div>
  )
}

export default App
