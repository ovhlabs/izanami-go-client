# izanami-go-client
Go client for [izanami](https://github.com/maif/izanami)


### Usage

```
c := New("host", "clientID", "clientSecret")
	
// List all features
features, errF := c.Feature().listAll()
if errF != nil {
	return errF
}

// Create a feature
f := FeatureModel{
	ID: "my-feature",
	Enabled: true,
	Strategy: NoStrategy,
}
if err := c.Feature().create(f); err != nil {
	return err
}

// Get a feature
myFeature, errF := c.Feature().get(f.ID)
if errF != nil {
	return errF
}

// Update a feature
if err := c.Feature().update(myFeature); err != nil {
	return err
}

// Delete a feature
if err := c.Feature().delete(myFeature.ID); err != nil {
	return err
}
```
