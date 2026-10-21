# IA04 Note d'intention
Rayan BARREDDINE, Benoît DE MISCAULT, Agathe RABASSE, Myriam REMILA

## Sujet : Optimisation multi-agents de la recharge de véhicules électriques sous contrainte de puissance du réseau éléctrique


## Descriptif
 
Dans ses mesures d'électrification annoncées en 2026, le gouvernement français a donné pour objectif que 2/3 des véhicules neufs soient électriques. Toutefois, l'éléctrification massive soulève de nouveaux enjeux quand à la capacité du réseau électrique à subvenir aux besoins de ce nouveau parc automobile. En particulier, la recharge simultanée de plusieurs véhicules électriques peut dépasser la puissance locale (par exemple à l'échelle d'un parking) du réseau électrique.

Notre simulation modélise ce site comme un système multi-agents où la puissance (en kW) doit être répartie entre tous les véhicules branchés, qui n'ont ni les mêmes besoins ni les mêmes contraintes.

### Agents

**Véhicules** : Chacun est un agent autonome (une goroutine) défini par son heure d'arrivée, son heure de départ prévue, l'énergie à recevoir, sa puissance de charge maximale. Le véhicule représente son conducteur avec ses contraintes. 
Les véhicules devront partager une puissance totale limitée.


**Coordinateur** : Il pourra répartir la puissance entre les véhicules selon les stratégies que nous choisirons d'étudier, par exemple : partage équitable, priorité aux véhicules qui doivent bientôt repartir ou priorité aux batteries les plus faibles.


## Question à laquelle notre simulation répond

Quelle stratégie de répartition de la puissance permet de répondre le plus efficacement aux besoins de recharge des véhicules avant leur départ, tout en respectant la limite de puissance du site ?

## Pistes de réflexions

- Utilisation d'un LLM pour rendre le comportement des usagers plus "humain".

## Indicateurs clés

- **Respect de la contrainte de puissance** : vérification que la puissance totale allouée ne dépasse jamais la puissance disponible
- **Taux de véhicules prêts au départ** : proportion des véhicules ayant atteint leur niveau de charge cible avant l’heure de départ
- **Énergie totale distribuée** : quantité d’énergie délivrée pendant la simulation
- **Durée de recharge** : temps nécessaire pour atteindre la cible
- **Délai de recharge** : temps d'attente avant le début de la charge
- **Satisfaction de l'usager** : Energie non fournie kWh manquants au départ pour véhicules non satisfaits en tenant compte de l'impact que cela peut avoir sur les contraintes de l'usager. 


## Sources 
https://www.info.gouv.fr/actualite/electrification-les-mesures-annoncees-par-le-gouvernement

