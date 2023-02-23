# Beast hunter

Beast hunter is a CLI game in which Natelus fights the beast on the basis of the stats

## Run Locally

Start docker

```bash
    docker build -t beast-hunter .
```

```bash
    docker run beast-hunter
```

Run `beast-hunter` code

```bash
  go run cmd/beast-hunter/main.go
```

Create a `beast-hunter` build

```bash
  go build -o build/ ./...
```

To run `beast-hunter` tests, run the following command

```bash
  go run test ./...
```

## Run Locally using make

Start docker

```bash
  make start-docker
```

Run `beast-hunter` code

```bash
  make run
```

Create a `beast-hunter` build

```bash
  make build
```

Lint `beast-hunter` code

```bash
  make lint
```

To run `beast-hunter` tests, run the following command

```bash
  make test
```

## Configuration variables are loaded in the following order

1. .default.env
2. .$env.env
3. environment variables

## Environment Variables

| ENV_VAR                            | Type        | Description                          |
| :--------------------------------- | :---------- | :----------------------------------- |
| `HERO_HEALTH_MAX`                  | `int`       | Maximum Hero Health                  |
| `HERO_HEALTH_MIN`                  | `int`       | Minimum Hero Health                  |
| `HERO_STRENGTH_MAX`                | `int`       | Maximum Hero Strength                |
| `HERO_STRENGTH_MIN`                | `int`       | Minimum Hero Strength                |
| `HERO_DEFENCE_MAX`                 | `int`       | Maximum Hero Defence                 |
| `HERO_DEFENCE_MIN`                 | `int`       | Minimum Hero Defence                 |
| `HERO_SPEED_MAX`                   | `int`       | Maximum Hero Speed                   |
| `HERO_SPEED_MIN`                   | `int`       | Minimum Hero Speed                   |
| `HERO_LUCK_MAX`                    | `int`       | Maximum Hero Luck                    |
| `HERO_LUCK_MIN`                    | `int`       | Minimum Hero Luck                    |
| `BEAST_HEALTH_MAX`                 | `int`       | Maximum Beast Health                 |
| `BEAST_HEALTH_MIN`                 | `int`       | Minimum Beast Health                 |
| `BEAST_STRENGTH_MAX`               | `int`       | Maximum Beast Strength               |
| `BEAST_STRENGTH_MIN`               | `int`       | Minimum Beast Strength               |
| `BEAST_DEFENCE_MAX`                | `int`       | Maximum Beast Defence                |
| `BEAST_DEFENCE_MIN`                | `int`       | Minimum Beast Defence                |
| `BEAST_SPEED_MAX`                  | `int`       | Maximum Beast Speed                  |
| `BEAST_SPEED_MIN`                  | `int`       | Minimum Beast Speed                  |
| `BEAST_LUCK_MAX`                   | `int`       | Maximum Beast Luck                   |
| `BEAST_LUCK_MIN`                   | `int`       | Minimum Beast Luck                   |
| `HERO_SKILL_LIST`                  | `[]string`  | Names of the hero skills             |
| `HERO_DAMAGE_MODIFIER_LIST`        | `[]float64` | Damage modifier for each skill       |
| `HERO_STRIKE_COUNT_MODIFIER_LIST`  | `[]float64` | Strike count modifier for each skill |
| `HERO_SKILL_CHANCE_LIST`           | `[]float64` | Chance for each skill                |
| `BEAST_SKILL_LIST`                 | `[]string`  | Names of the beast skills            |
| `BEAST_DAMAGE_MODIFIER_LIST`       | `[]float64` | Damage modifier for each skill       |
| `BEAST_STRIKE_COUNT_MODIFIER_LIST` | `[]float64` | Strike count modifier for each skill |
| `BEAST_SKILL_CHANCE_LIST`          | `[]float64` | Chance for each skill                |

## Character Statistics

#### Hero Stats

| Stat       | Type  | Description                   |
| :--------- | :---- | :---------------------------- |
| `Health`   | `Int` | Hero Health                   |
| `Strength` | `Int` | Hero Strength while attacking |
| `Defence`  | `Int` | Ability to defend the attack  |
| `Speed`    | `Int` | Ability to attack first       |
| `Luck`     | `Int` | Ability to dodge the attack   |

#### Skills

| Skill          | Description                                                                 |
| :------------- | :-------------------------------------------------------------------------- |
| `Rapid Strike` | Strike twice while it's his turn to attack (with 10% chance)                |
| `Magic Sheild` | Takes only half of the usual damage when an enemy attacks (with 20% chance) |

#### Beast Stats

| Stat       | Type  | Description                    |
| :--------- | :---- | :----------------------------- |
| `Health`   | `Int` | Beast Health                   |
| `Strength` | `Int` | Beast Strength while attacking |
| `Defence`  | `Int` | Ability to defend the attack   |
| `Speed`    | `Int` | Ability to attack first        |
| `Luck`     | `Int` | Ability to dodge the attack    |

### NOTE

<b>Using the configuration we can add skills to any character and also introduce different modifiers for different stats of the character. <b>
