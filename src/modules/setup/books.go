package setup

const updateAptCache = `---
name: updateAptCache
target_os: "Linux"
variables:
  - name: pre
    value: "DEBIAN_FRONTEND=noninteractive timeout 300"
tasks:
  - name: apt-get update
    command: "{{ pre }} apt-get update -y"
  - name: apt list
    command: "{{ pre }} apt list --upgradeable 2>/dev/null | tail -n +2"
  - name: apt list amount
    command: "{{ pre }} apt list --upgradable 2>/dev/null | tail -n +2 | wc -l"
`

const updateOs = `---
name: updateOs
target_os: "Linux"
variables:
  - name: pre
    value: "DEBIAN_FRONTEND=noninteractive timeout 300"
tasks:
  - name: dpkg configure
    command: "{{ pre }} dpkg --configure -a --force-confold"
  - name: apt-get update
    command: "{{ pre }} apt-get update"
  - name: apt-get fix broken install
    command: "{{ pre }} apt-get --fix-broken install -y"
  - name: apt-get upgrade
    command: >-
      {{ pre }} apt-get -q
      -o Dpkg::Options::='--force-confold'
      upgrade -y
  - name: apt-get autoremove
    command: "{{ pre }} apt-get -q autoremove -y"
`
