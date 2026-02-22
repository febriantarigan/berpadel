import { useState } from 'react'

const TOURNAMENT_TYPES = [
  { value: 'americano', label: 'Americano', desc: 'Rotating partners' },
  { value: 'americano_fixed', label: 'Americano Fixed', desc: 'Fixed teams' },
  { value: 'mix_americano', label: 'Mix Americano', desc: 'Mixed, rotating' },
  { value: 'mix_americano_fixed', label: 'Mix Americano Fixed', desc: 'Mixed, fixed teams' },
  { value: 'mexicano', label: 'Mexicano', desc: 'Mexicano format' },
] as const

const STATUS_OPTIONS = [
  { value: 'draft', label: 'Draft' },
  { value: 'ongoing', label: 'Ongoing' },
  { value: 'completed', label: 'Completed' },
]

type Props = {
  open: boolean
  onClose: () => void
  onSubmit: (data: {
    name: string
    type: string
    status: string
    location: string
    season: string
    courts: string[]
    maxPoints: number
    players: { name: string; gender: string }[]
  }) => void | Promise<void>
  error: string | null
}

export default function CreateTournamentModal({ open, onClose, onSubmit, error }: Props) {
  const [name, setName] = useState('')
  const [type, setType] = useState('americano')
  const [status, setStatus] = useState('draft')
  const [location, setLocation] = useState('')
  const [season, setSeason] = useState(new Date().getFullYear().toString())
  const [maxPoints, setMaxPoints] = useState(15)
  const [courts, setCourts] = useState('Court 1')
  const [players, setPlayers] = useState<{ name: string; gender: string }[]>([
    { name: '', gender: 'male' },
    { name: '', gender: 'male' },
  ])
  const [submitting, setSubmitting] = useState(false)

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    const validPlayers = players.filter((p) => p.name.trim())
    if (validPlayers.length < 4) {
      return
    }
    setSubmitting(true)
    await onSubmit({
      name: name.trim(),
      type,
      status,
      location: location.trim(),
      season,
      courts: courts.split(',').map((c) => c.trim()).filter(Boolean),
      maxPoints,
      players: validPlayers,
    })
    setSubmitting(false)
    if (!error) {
      setName('')
      setLocation('')
      setPlayers([{ name: '', gender: 'male' }, { name: '', gender: 'male' }])
    }
  }

  const addPlayer = () => setPlayers((p) => [...p, { name: '', gender: 'male' }])
  const removePlayer = (i: number) => setPlayers((p) => p.filter((_, idx) => idx !== i))
  const updatePlayer = (i: number, field: 'name' | 'gender', value: string) => {
    setPlayers((p) => p.map((pl, idx) => (idx === i ? { ...pl, [field]: value } : pl)))
  }

  const validPlayers = players.filter((p) => p.name.trim())
  const canSubmit = name.trim() && validPlayers.length >= 4

  if (!open) return null

  return (
    <dialog open className="modal modal-open">
      <div className="modal-box max-w-lg max-h-[90vh] overflow-y-auto">
        <h3 className="font-bold text-lg mb-4">Create Tournament</h3>
        <form onSubmit={handleSubmit} className="space-y-4">
          <div>
            <label className="label">
              <span className="label-text">Name</span>
            </label>
            <input
              type="text"
              className="input input-bordered w-full"
              placeholder="e.g. Friday Night Mix"
              value={name}
              onChange={(e) => setName(e.target.value)}
              required
            />
          </div>

          <div>
            <label className="label">
              <span className="label-text">Format</span>
            </label>
            <select
              className="select select-bordered w-full"
              value={type}
              onChange={(e) => setType(e.target.value)}
            >
              {TOURNAMENT_TYPES.map((t) => (
                <option key={t.value} value={t.value}>
                  {t.label} — {t.desc}
                </option>
              ))}
            </select>
          </div>

          <div>
            <label className="label">
              <span className="label-text">Status</span>
            </label>
            <select
              className="select select-bordered w-full"
              value={status}
              onChange={(e) => setStatus(e.target.value)}
            >
              {STATUS_OPTIONS.map((s) => (
                <option key={s.value} value={s.value}>
                  {s.label}
                </option>
              ))}
            </select>
          </div>

          <div className="grid grid-cols-2 gap-4">
            <div>
              <label className="label">
                <span className="label-text">Location</span>
              </label>
              <input
                type="text"
                className="input input-bordered w-full"
                placeholder="e.g. Baerpadel Berlin"
                value={location}
                onChange={(e) => setLocation(e.target.value)}
              />
            </div>
            <div>
              <label className="label">
                <span className="label-text">Season</span>
              </label>
              <input
                type="text"
                className="input input-bordered w-full"
                placeholder="2025"
                value={season}
                onChange={(e) => setSeason(e.target.value)}
              />
            </div>
          </div>

          <div className="grid grid-cols-2 gap-4">
            <div>
              <label className="label">
                <span className="label-text">Max points per set</span>
              </label>
              <input
                type="number"
                min={1}
                className="input input-bordered w-full"
                value={maxPoints}
                onChange={(e) => setMaxPoints(Number(e.target.value) || 15)}
              />
            </div>
            <div>
              <label className="label">
                <span className="label-text">Courts (comma-separated)</span>
              </label>
              <input
                type="text"
                className="input input-bordered w-full"
                placeholder="Court 1, Court 2"
                value={courts}
                onChange={(e) => setCourts(e.target.value)}
              />
            </div>
          </div>

          <div>
            <div className="label flex justify-between items-center">
              <span className="label-text">Players (min 4)</span>
              <button type="button" className="btn btn-ghost btn-xs" onClick={addPlayer}>
                + Add
              </button>
            </div>
            <div className="space-y-2 max-h-48 overflow-y-auto">
              {players.map((p, i) => (
                <div key={i} className="flex gap-2">
                  <input
                    type="text"
                    className="input input-bordered flex-1 input-sm"
                    placeholder="Player name"
                    value={p.name}
                    onChange={(e) => updatePlayer(i, 'name', e.target.value)}
                  />
                  <select
                    className="select select-bordered select-sm w-24"
                    value={p.gender}
                    onChange={(e) => updatePlayer(i, 'gender', e.target.value)}
                  >
                    <option value="male">M</option>
                    <option value="female">F</option>
                  </select>
                  <button
                    type="button"
                    className="btn btn-ghost btn-sm btn-square"
                    onClick={() => removePlayer(i)}
                    disabled={players.length <= 2}
                  >
                    ×
                  </button>
                </div>
              ))}
            </div>
            {validPlayers.length > 0 && validPlayers.length < 4 && (
              <p className="text-error text-sm mt-1">Add at least {4 - validPlayers.length} more player(s)</p>
            )}
          </div>

          {error && (
            <div className="alert alert-error">
              <span>{error}</span>
            </div>
          )}

          <div className="modal-action">
            <button type="button" className="btn btn-ghost" onClick={onClose} disabled={submitting}>
              Cancel
            </button>
            <button type="submit" className="btn btn-primary" disabled={!canSubmit || submitting}>
              {submitting ? 'Creating…' : 'Create'}
            </button>
          </div>
        </form>
      </div>
      <form method="dialog" className="modal-backdrop" onClick={onClose}>
        <button type="button">close</button>
      </form>
    </dialog>
  )
}
