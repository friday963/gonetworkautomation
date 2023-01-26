package templatefiles

const CeosTemplate = `
!
configure
!
ip routing
!
{{range $vrf_routing := .Vrfs}}
ip routing vrf {{$vrf_routing.Name}}
!
{{end}}

{{range $logicalinterfaces := .Interfaces.Logical}}
interface {{$logicalinterfaces.Name}}
ip address {{ $logicalinterfaces.IP }}/32
!
{{end}}
{{range $interfaces := .Interfaces.Physical}}
interface {{$interfaces.Name}}
ip address {{$interfaces.IP}}
{{end}}
!

{{range $vlans := .Vlans}}
vlan {{$vlans.Number}}
  name {{$vlans.Name}}
{{end}}
!
{{range $prefixlist := .PrefixLists}}
ip prefix-list {{$prefixlist.Name}} {{$prefixlist.Seq}} {{$prefixlist.PermitDeny}} {{$prefixlist.Prefix}} {{$prefixlist.Mask}} {{if and $prefixlist.LeGe $prefixlist.LeGeMask }} {{$prefixlist.LeGe}} {{$prefixlist.LeGeMask}}{{else}} \n {{end}}
!
{{end}}
`