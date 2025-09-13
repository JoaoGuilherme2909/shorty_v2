package store

import "context"

type Store struct{ Client *Client }

func (rds *Store) SetUrl(ctx context.Context, key, value string) error {
	err := rds.Client.Client.Set(ctx, key, value, 0).Err()

	if err != nil {
		return err
	}

	return nil
}

func (rds *Store) GetUrl(ctx context.Context, key string) (string, error) {
	value, err := rds.Client.Client.Get(ctx, key).Result()

	if err != nil {
		return "", err
	}

	return value, nil
}
