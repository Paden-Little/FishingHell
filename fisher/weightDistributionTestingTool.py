import random
import math
import matplotlib.pyplot as plt

def custom_pdf(x, peak, spread):
    return math.exp(-((x - peak) ** 2) / (2 * spread ** 2))


def plot_fish_weight_distribution(low_weight, high_weight, peak, stddev, num_samples=10000):
    weights = [generate_fish_weight_custom(low_weight, high_weight, peak, stddev) for _ in range(num_samples)]
    plt.hist(weights, bins=50, density=True, alpha=0.6, color='skyblue', edgecolor='black')
    plt.title(f'Fish Weight Distribution (stddev_factor={stddev}), (peak={peak})')
    plt.xlabel('Weight (ounces)')
    plt.ylabel('Probability Density')
    plt.axvline(x=(low_weight + high_weight) / 2, color='red', linestyle='dashed', linewidth=1, label='Mean')
    plt.legend()
    plt.show()


def generate_fish_weight_custom(low_weight, high_weight, peak=None, stddev=None, num_samples=100):
    if peak is None:
        peak = (low_weight + high_weight) / 2
    if stddev is None:
        spread = (high_weight - low_weight) / 6
    else:
        spread = (high_weight - low_weight) / stddev

    weights = []
    for _ in range(num_samples):
        weight = random.uniform(low_weight, high_weight)
        probability = custom_pdf(weight, peak, spread)
        if random.random() < probability:
            weights.append(weight)

    if weights:
        return round(random.choice(weights), 2)
    else:
        # Retry if no weights were accepted
        return generate_fish_weight_custom(low_weight, high_weight, peak, spread, num_samples)

low_weight = 300
high_weight = 625
peak = 320
stddev = 100


plot_fish_weight_distribution(low_weight, high_weight, peak, stddev)

