const status = document.getElementById('status');
const rows = document.getElementById('rows');

const chart = new Chart(document.getElementById('chart'), {
  type: 'line',
  data: {
    labels: [],
    datasets: [
      { label: 'Limite du site (kW)', data: [], borderColor: 'crimson', borderDash: [6, 4], pointRadius: 0 },
      { label: 'Puissance allouée (kW)', data: [], borderColor: 'steelblue', pointRadius: 0 },
    ],
  },
  options: { animation: false, scales: { y: { beginAtZero: true, title: { display: true, text: 'kW' } } } },
});

async function main() {
  const cfg = await (await fetch('/config')).json();
  status.textContent = `Stratégie ${cfg.Strategy} · ${cfg.Ticks} ticks de ${cfg.MinutesPerTick} min`;

  const es = new EventSource('/events');
  es.onmessage = (e) => {
    const s = JSON.parse(e.data);

    chart.data.labels.push(`${Math.floor(s.T * cfg.MinutesPerTick / 60)}h${String(s.T * cfg.MinutesPerTick % 60).padStart(2, '0')}`);
    chart.data.datasets[0].data.push(s.PSite);
    chart.data.datasets[1].data.push(s.TotalAllocated);
    chart.update();

    rows.innerHTML = (s.Vehicles || []).map(v =>
      `<tr><td>${v.ID}</td><td>${v.State}</td><td>${(v.Soc * 100).toFixed(0)} %</td>` +
      `<td>${v.AllocatedKW.toFixed(1)}</td><td>${v.EnergyNeeded.toFixed(1)}</td></tr>`
    ).join('');

    if (s.T >= cfg.Ticks - 1) { es.close(); status.textContent += ' · terminé'; }
  };
  es.onerror = () => { status.textContent = 'Connexion perdue'; };
}

main();