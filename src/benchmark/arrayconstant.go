package benchmark

type status string

func (s status) toString() string {
	return string(s)
}

const (
	ok                          status = "OK"
	routingGatewayFallback      status = "ROUTING_GATEWAY_FALLBACK"
	googleComputeRoutesFallback status = "GOOGLE_COMPUTEROUTES_FALLBACK"
	graphhopperFallback         status = "GRAPHHOPPER_FALLBACK"
	googleFallback              status = "GOOGLE_FALLBACK"
)

var validStatuses = [5]status{
	ok,
	routingGatewayFallback,
	googleComputeRoutesFallback,
	graphhopperFallback,
	googleFallback,
}

func GetValidStatus() []status {
	return []status{
		ok,
		routingGatewayFallback,
		googleComputeRoutesFallback,
		graphhopperFallback,
		googleFallback}
}

func CallMethod1() {

	isStatusValid1(googleComputeRoutesFallback)

}

func CallMethod2() {

	isStatusValid2(googleComputeRoutesFallback)

}

func CallMethod3() {

	isStatusValid3(googleComputeRoutesFallback)

}

func isStatusValid1(returnedStatus status) bool {
	validStatuses := []status{
		ok,
		routingGatewayFallback,
		googleComputeRoutesFallback,
		graphhopperFallback,
		googleFallback,
	}

	for _, s := range validStatuses {
		if returnedStatus == s {
			return true
		}
	}

	return false
}

func isStatusValid2(returnedStatus status) bool {

	for _, s := range validStatuses {
		if returnedStatus == s {
			return true
		}
	}
	return false
}

func isStatusValid3(returnedStatus status) bool {

	for _, s := range GetValidStatus() {
		if returnedStatus == s {
			return true
		}
	}
	return false
}
